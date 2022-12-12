package controllers

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
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

func listObjects(bucketName string) ([]string, error) {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var objectList []string
	objects := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		object, err := objects.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objectList = append(objectList, object.Name)
	}
	return objectList, nil
}

func listObjectsWithPrefix(bucketName, prefix string) ([]string, error) {
	ctx := gcs.GetStorageClientContext()
	client := gcs.GetStorageClient()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var objectList []string
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
		objectList = append(objectList, object.Name)
	}
	return objectList, nil
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

func getFileURL(bucketName, objectName string) (string, error) {
	fmt.Println(os.Getenv("SERVER_STORAGE_PEM_LOCATION"))
	pkey, err := ioutil.ReadFile(os.Getenv("SERVER_STORAGE_PEM_LOCATION"))
	if err != nil {
		return "", nil
	}

	fmt.Println(bucketName)
	fmt.Println(objectName)

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

	// ctx := gcs.GetStorageClientContext()
	// client := gcs.GetStorageClient()

	// ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	// defer cancel()

	// rc, err := client.Bucket(bucket).Object(object).NewReader(ctx)
	// if err != nil {
	// 	return err
	// }
	// defer rc.Close()

	// newFile, err := os.Create(os.Getenv("SERVER_STORAGE_LOCATION") + "/" + r)
	// if err != nil {
	// 	return err
	// }

	// if _, err := io.Copy(newFile, rc); err != nil {
	// 	return err
	// }

	// if err = f.Close(); err != nil {
	// 	return err
	// }
	// return  nil
}
