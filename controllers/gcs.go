package controllers

import (
	"context"
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