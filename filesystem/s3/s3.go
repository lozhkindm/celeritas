package s3

import (
	"github.com/lozhkindm/celeritas/filesystem"
)

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

func (s *S3) Put(filename, folder string) error {
	return nil
}

func (s *S3) Get(dst string, items ...string) error {
	return nil
}

func (s *S3) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	return entries, nil
}

func (s *S3) Delete(toDelete []string) (bool, error) {
	return true, nil
}
