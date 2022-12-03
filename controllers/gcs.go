package controllers

import (
	"context"
	"io"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/uet-class/uet-class-backend/config"
)

func createBucket(bucketName string) error {
	conf := config.GetConfig()
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	newBucket := &storage.BucketAttrs{
		Location:	conf.GetString("GCS_BUCKET_LOCATION"),
	}

	bucketHandle := client.Bucket(bucketName)
	if err := bucketHandle.Create(ctx, conf.GetString("GCP_PROJECT_ID"), newBucket); err != nil {
		return err
	}
	return nil
}

func uploadObject(bucketName string, file multipart.FileHeader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

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
