package storage

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/uet-class/uet-class-backend/config"
)

var storageClient *storage.Client
var ctx context.Context
var err error

func InitStorageClient() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", config.GetConfig().GetString("GOOGLE_APPLICATION_CREDENTIALS"))
	ctx = context.Background()

	storageClient, err = storage.NewClient(ctx)
	if err != nil {
		log.Fatal(http.StatusInternalServerError, err)
	}
	println("It is not dying...", storageClient.Buckets(ctx, "uet-class"))
	defer storageClient.Close()

	// Sets the name for the new bucket.
	bucketName := "my-new-bucket"

	// Creates a Bucket instance.
	bucket := storageClient.Bucket(bucketName)

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := bucket.Create(ctx, config.GetConfig().GetString("GOOGLE_PROJECT_ID"), nil); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", bucketName)
}

func GetStorageClient() *storage.Client {
	return storageClient
}

func GetStorageClientContext() context.Context {
	return ctx
}
