package s3

import (
	"github.com/lozhkindm/celeritas/filesystem"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
