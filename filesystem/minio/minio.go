package minio

import (
	"context"
	"fmt"
	"path"

	"github.com/lozhkindm/celeritas/filesystem"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Endpoint string
	Key      string
	Secret   string
	UseSSL   bool
	Region   string
	Bucket   string
}

func (m *Minio) getCredentials() (*minio.Client, error) {
	return minio.New(m.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.Key, m.Secret, ""),
		Secure: m.UseSSL,
	})
}

func (m *Minio) Put(filename, folder string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objectName := path.Base(filename)
	client, err := m.getCredentials()
	if err != nil {
		return err
	}
	_, err = client.FPutObject(ctx, m.Bucket, fmt.Sprintf("%s/%s", folder, objectName), filename, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Minio) Get(dst string, items ...string) error {
	//TODO implement me
	panic("implement me")
}

func (m *Minio) List(prefix string) ([]filesystem.ListEntry, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Minio) Delete(toDelete []string) bool {
	//TODO implement me
	panic("implement me")
}
