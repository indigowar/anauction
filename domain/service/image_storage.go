package service

import "context"

type ImageStorage interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Add(ctx context.Context, key string, value []byte) error
}
