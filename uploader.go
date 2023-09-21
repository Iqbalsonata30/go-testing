package gotesting

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

type GoogleStorageUploader interface {
	NewWriter(ctx context.Context) *storage.Writer
}

type Uploader struct {
	svc GoogleStorageUploader
}

func (u Uploader) Upload() error {
	ctx := context.Background()
	wc := u.svc.NewWriter(ctx)
	fmt.Println(wc)
	return nil
}
