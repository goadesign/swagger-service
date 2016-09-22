package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	sg "google.golang.org/api/storage/v1"
)

const (
	projectID  = "goa-swagger"
	bucketName = "artifacts.goa-swagger.appspot.com"
)

var (
	ctx    context.Context
	bucket *storage.BucketHandle
)

const (
	scope = sg.DevstorageReadWriteScope
)

func init() {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Unable to get storage client: %v", err)
	}
	bucket = client.Bucket(bucketName)

}

// Load attempts to load the swagger spec for the given package and given revision SHA.
// It returns the swagger spec content and true on success, nil and false if not found.
func Load(sha string) ([]byte, error) {
	object := bucket.Object(ObjectName(sha))
	rc, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return ioutil.ReadAll(rc)
}

// Save saves the given swagger spec to the cache.
func Save(b []byte, sha string) error {
	object := bucket.Object(ObjectName(sha))
	wc := object.NewWriter(ctx)
	defer wc.Close()
	wc.ContentType = "text/plain"
	wc.ACL = []storage.ACLRule{{storage.AllAuthenticatedUsers, storage.RoleOwner}}
	_, err := wc.Write(b)
	return err
}

// ObjectName returns the cloud storage object name for the given SHA.
func ObjectName(sha string) string {
	return fmt.Sprintf("specs/%s", sha)
}
