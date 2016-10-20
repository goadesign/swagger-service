package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/swagger-service/app"
)

// SpecController implements the spec resource.
type SpecController struct {
	*goa.Controller
}

// NewSpecController creates a spec controller.
func NewSpecController(service *goa.Service) *SpecController {
	return &SpecController{Controller: service.NewController("Spec")}
}

// Show clones the remote repo, runs "goagen swagger" and returns the corresponding JSON.
// It uses cloud storage to cache the JSON using the git commit SHA in the object name.
func (c *SpecController) Show(ctx *app.ShowSpecContext) error {
	logger := goalog15.Logger(ctx)
	tmpGoPath, err := ioutil.TempDir("", "swagger-service-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpGoPath)
	packagePath := strings.TrimPrefix(ctx.PackagePath, "/")
	elems := strings.Split(packagePath, "/")
	if len(elems) < 3 {
		return fmt.Errorf("invalid package path %s", packagePath)
	}
	var branch string
	parts := strings.Split(elems[len(elems)-1], "@")
	design := parts[0]
	if len(parts) > 1 {
		branch = parts[1]
		elems[len(elems)-1] = design
		packagePath = strings.Join(elems, "/")
	}
	repo := strings.Join(elems[:3], "/")
	dir := strings.Join(elems[:2], "/")
	dir = filepath.Join(tmpGoPath, "src", dir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	sha, err := clone("https://"+repo, dir, branch)
	if err != nil {
		return ctx.UnprocessableEntity([]byte(fmt.Sprintf("git clone: %s", err.Error())))
	}
	b, err := Load(sha)
	if err != nil {
		logger.Info("cache miss", "sha", sha)
	} else {
		return ctx.OK(b)
	}
	genCmd := exec.Command("goagen", "-o", tmpGoPath, "swagger", "-d", packagePath)
	genCmd.Env = []string{
		fmt.Sprintf("GOPATH=%s:%s", tmpGoPath, os.Getenv("GOPATH")),
		"PATH=" + os.Getenv("PATH"),
	}
	out, err := genCmd.CombinedOutput()
	if err != nil {
		if len(out) == 0 {
			out = []byte(err.Error())
		}
		return ctx.UnprocessableEntity(out)
	}
	b, err = ioutil.ReadFile(filepath.Join(tmpGoPath, "swagger", "swagger.json"))
	if err != nil {
		return ctx.UnprocessableEntity([]byte(err.Error()))
	}
	if sha != "" {
		err := Save(b, sha)
		if err != nil {
			logger.Error("failed to save swagger spec", "package", packagePath, "error", err.Error())
		}
	}
	return ctx.OK(b)
}

// clone does a shallow clone of the repo in the given directory and return the SHA
// If there is no branch specified, try "go1" branch followed by "master" branch.
// If the branch is not available return empty SHA with error
func clone(repo, tmpDir, newbranch string) (string, error) {
	var branch string
	shallowClone := func() error {
		gitCmd := exec.Command("git", "clone", "--depth=1", "--single-branch", "--branch", branch, repo)
		gitCmd.Dir = tmpDir
		return gitCmd.Run()
	}

	if newbranch == "" {
		branch = "go1"
		if err := shallowClone(); err != nil {
			branch = "master"
			if err = shallowClone(); err != nil {
				return "", fmt.Errorf("failed to clone %s", repo)
			}
		}

	} else {
		branch = newbranch
		if err := shallowClone(); err != nil {
			return "", fmt.Errorf("failed to clone %s branch: %s", repo, branch)
		}
	}

	gitCmd := exec.Command("git", "rev-parse", branch)
	gitCmd.Dir = filepath.Join(tmpDir, filepath.Base(repo))
	out, err := gitCmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
