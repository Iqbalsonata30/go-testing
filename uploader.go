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
	bucket := "bucket-name"
	object := "object-name"

	ctx := context.Background()
	str := new(storage.Client)
	o := str.Bucket(bucket).Object(object)
	o = o.If(storage.Conditions{DoesNotExist: true})

	wc := o.NewWriter(ctx)
	fmt.Println(wc)
	return nil
}
