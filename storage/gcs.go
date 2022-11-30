package storage

import (
	"context"
	"log"
	"net/http"
	"os"

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
}

func GetStorageClient() *storage.Client {
	return storageClient
}

func GetStorageClientContext() context.Context {
	return ctx
}
