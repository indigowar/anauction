package minio

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"

	"github.com/indigowar/anauction/domain/service"
)

const bucketName = "images"

type ImageStorage struct {
	client *minio.Client
}

var _ service.ImageStorage = &ImageStorage{}

// Add implements service.ImageStorage.
func (i *ImageStorage) Add(ctx context.Context, key string, value []byte) error {
	reader := bytes.NewReader(value)

	_, err := i.client.PutObject(ctx, bucketName, key, reader, reader.Size(), minio.PutObjectOptions{
		ContentType: "image/png",
	})
	if err != nil {
		// TODO: handle this error properly
		return err
	}

	return nil
}

// Get implements service.ImageStorage.
func (i *ImageStorage) Get(ctx context.Context, key string) ([]byte, error) {
	object, err := i.client.GetObject(ctx, bucketName, key, minio.GetObjectOptions{})
	if err != nil {
		// TODO: handle this error properly
		return nil, err
	}

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, object); err != nil {
		// TODO: handle this error properly
		return nil, err
	}

	return buf.Bytes(), nil
}

func NewImageStorage(client *minio.Client) (*ImageStorage, error) {
	if err := createBucketIfDoesNotExist(client, bucketName); err != nil {
		return nil, err
	}

	return &ImageStorage{
		client: client,
	}, nil
}
