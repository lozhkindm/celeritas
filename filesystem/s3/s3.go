package s3

import (
	"bytes"
	"net/http"
	"os"
	"path"

	"github.com/lozhkindm/celeritas/filesystem"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	Key      string
	Secret   string
	Region   string
	Endpoint string
	Bucket   string
}

func (s *S3) Put(filename, folder string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	fileinfo, err := file.Stat()
	if err != nil {
		return err
	}
	bts := make([]byte, fileinfo.Size())
	if _, err := file.Read(bts); err != nil {
		return err
	}
	reader := bytes.NewReader(bts)
	filetype := http.DetectContentType(bts)
	client := credentials.NewStaticCredentials(s.Key, s.Secret, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(s.Endpoint),
		Region:      aws.String(s.Region),
		Credentials: client,
	}))
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(path.Join(folder, path.Base(filename))),
		Body:        reader,
		ACL:         aws.String("public-read"),
		ContentType: aws.String(filetype),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3) Get(dst string, items ...string) error {
	return nil
}

func (s *S3) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	client := credentials.NewStaticCredentials(s.Key, s.Secret, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint:    aws.String(s.Endpoint),
		Region:      aws.String(s.Region),
		Credentials: client,
	}))
	service := s3.New(sess)
	input := &s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(prefix),
	}
	result, err := service.ListObjects(input)
	if err != nil {
		return nil, err
	}
	for _, content := range result.Contents {
		entries = append(entries, filesystem.ListEntry{
			Etag:         *content.ETag,
			LastModified: *content.LastModified,
			Key:          *content.Key,
			Size:         float64(*content.Size) / 1024 / 1024,
		})
	}
	return entries, nil
}

func (s *S3) Delete(toDelete []string) (bool, error) {
	return true, nil
}
