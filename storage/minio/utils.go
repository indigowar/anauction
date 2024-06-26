package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
)

func createBucketIfDoesNotExist(client *minio.Client, bucketName string) error {
	ctx := context.Background()

	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		return client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	}

	return nil
}
