package gcs

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

var storageClient *storage.Client
var ctx context.Context
var err error

func InitStorageClient() {
	ctx = context.Background()

	serviceAccount := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(serviceAccount))
	if err != nil {
		log.Fatal(http.StatusInternalServerError, err.Error())
	}
}

func GetStorageClient() *storage.Client {
	return storageClient
}

func GetStorageClientContext() context.Context {
	return ctx
}
