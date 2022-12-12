package controllers

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/uet-class/uet-class-backend/gcs"
	"google.golang.org/api/iterator"
)

func createBucket(bucketName string) error {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	newBucket := &storage.BucketAttrs{
		Location: os.Getenv("GCS_BUCKET_LOCATION"),
	}

	bucketHandle := client.Bucket(bucketName)
	if err := bucketHandle.Create(ctx, os.Getenv("GCP_PROJECT_ID"), newBucket); err != nil {
		return err
	}
	return nil
}

func uploadObject(bucketName string, file multipart.FileHeader) error {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	uploadFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	objectHandle := client.Bucket(bucketName).Object(file.Filename)
	objectHandle = objectHandle.If(storage.Conditions{DoesNotExist: true})

	objectWriter := objectHandle.NewWriter(ctx)
	if _, err = io.Copy(objectWriter, uploadFile); err != nil {
		return err
	}

	if err := objectWriter.Close(); err != nil {
		return err
	}
	return nil
}

func listObjects(bucketName string) ([]interface{}, error) {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var files []interface{}
	objects := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		object, err := objects.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		file := map[string]string{
			"fileName":  object.Name,
			"createdAt": object.Created.String(),
			"prefix":    object.Prefix,
		}
		files = append(files, file)
	}
	return files, nil
}

func listObjectsWithPrefix(bucketName, prefix string) ([]interface{}, error) {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var files []interface{}
	objects := client.Bucket(bucketName).Objects(ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		object, err := objects.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		file := map[string]string{
			"fileName":  object.Name,
			"createdAt": object.Created.String(),
			"prefix":    object.Prefix,
		}
		files = append(files, file)
	}
	return files, nil
}

func deleteObject(bucketName string, objectName string) error {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	objectHandle := client.Bucket(bucketName).Object(objectName)

	objectAttributes, err := objectHandle.Attrs(ctx)
	if err != nil {
		return err
	}
	objectHandle = objectHandle.If(storage.Conditions{GenerationMatch: objectAttributes.Generation})

	if err := objectHandle.Delete(ctx); err != nil {
		return err
	}
	return nil
}

func getFileSignedURL(bucketName, objectName string) (string, error) {
	pkey, err := os.ReadFile(os.Getenv("SERVER_STORAGE_PEM_LOCATION"))
	if err != nil {
		return "", nil
	}

	url, err := storage.SignedURL(bucketName, objectName, &storage.SignedURLOptions{
		GoogleAccessID: "uc-backend-sa@uet-class.iam.gserviceaccount.com",
		PrivateKey:     pkey,
		Method:         "GET",
		Expires:        time.Now().Add(48 * time.Hour),
	})
	if err != nil {
		return "", err
	}
	return url, nil
}
