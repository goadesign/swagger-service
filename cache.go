package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	sg "google.golang.org/api/storage/v1"
	"google.golang.org/cloud"
	"google.golang.org/cloud/storage"
)

var (
	projectID  = "goa-swagger"
	bucketName = "artifacts.goa-swagger.appspot.com"
	bucket     *storage.Bucket
	ctx        context.Context
)

const (
	scope = sg.DevstorageReadWriteScope
)

func init() {
	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}
	ctx = cloud.NewContext(projectID, client)
	if _, err := storage.BucketInfo(ctx, bucketName); err != nil {
		log.Fatalf("Failed retrieving bucket information: %s", err.Error())
	}
}

// Load attempts to load the swagger spec for the given package and given revision SHA.
// It returns the swagger spec content and true on success, nil and false if not found.
func Load(packagePath, sha string) ([]byte, error) {
	rc, err := storage.NewReader(ctx, bucketName, ObjectName(sha))
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// Save saves the given swagger spec to the cache.
func Save(b []byte, packagePath, sha string) error {
	wc := storage.NewWriter(ctx, bucketName, ObjectName(sha))
	wc.ContentType = "text/plain"
	wc.ACL = []storage.ACLRule{{storage.AllAuthenticatedUsers, storage.RoleOwner}}
	_, err := wc.Write(b)
	return err
}

// ObjectName returns the cloud storage object name for the given SHA.
func ObjectName(sha string) string {
	return fmt.Sprintf("specs/%s", sha)
}
