package model

import "io"

type ProductUploadInput struct {
	Object      io.Reader
	ObjectName  string
	ObjectSize  int64
	BucketName  string
	ContentType string
}
