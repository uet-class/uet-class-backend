package storage

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
)

var storageClient *storage.Client
var ctx context.Context
var err error

func InitStorageClient() {
	ctx = context.Background()

	storageClient, err = storage.NewClient(ctx)
	if err != nil {
		log.Fatal(http.StatusInternalServerError, err)
	}
}

func GetStorageClient() *storage.Client {
	return storageClient
}

func GetStorageClientContext() context.Context {
	return ctx
}
