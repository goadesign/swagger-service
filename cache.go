package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

var (
	projectID  = "goa-swagger"
	bucketName = "artifacts.goa-swagger.appspot.com"
	service    *storage.Service
)

const (
	scope = storage.DevstorageReadWriteScope
)

func init() {
	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}
	service, err = storage.New(client)
	if err != nil {
		log.Fatalf("Unable to create storage service: %v", err)
	}
	if _, err := service.Buckets.Get(bucketName).Do(); err != nil {
		// Create a bucket.
		if res, err := service.Buckets.Insert(projectID, &storage.Bucket{Name: bucketName}).Do(); err == nil {
			fmt.Printf("Created bucket %v at location %v\n\n", res.Name, res.SelfLink)
		} else {
			log.Fatalf("Failed creating bucket %s: %v", bucketName, err)
		}
	}
}

// Load attempts to load the swagger spec for the given package and given revision SHA.
// It returns the swagger spec content and true on success, nil and false if not found.
func Load(packagePath, sha string) ([]byte, error) {
	objectName := "specs/" + sha
	res, err := service.Objects.Get(bucketName, objectName).Do()
	if err != nil {
		return nil, err
	}
	return res.MarshalJSON()
}

// Save saves the given swagger spec to the cache.
func Save(b []byte, packagePath, sha string) error {
	object := &storage.Object{Name: "specs/" + sha}
	buf := bytes.NewBuffer(b)
	_, err := service.Objects.Insert(bucketName, object).Media(buf).Do()
	return err
}
