package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/storage"
)

func main() {
	client, err := storage.NewClient(os.Getenv("STORAGE_ACCOUNT"), os.Getenv("STORAGE_ACCOUNT_KEY"), storage.DefaultBaseURL, storage.DefaultAPIVersion, true)
	if err != nil {
		log.Fatalf("Failed to create storage client: %s\n", err)
	}

	blobClient := client.GetBlobService()
	cnt := blobClient.GetContainerReference(os.Getenv("STORAGE_CONTAINER"))

	blobs, err := cnt.ListBlobs(storage.ListBlobsParameters{
		Include: &storage.IncludeBlobDataset{
			Snapshots: true,
			Copy:      true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to list blobs: %s\n", err)
	}

	for _, blob := range blobs.Blobs {
		fmt.Println(blob.Name)
	}
}
