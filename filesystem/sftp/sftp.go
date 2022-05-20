package sftp

import (
	"github.com/lozhkindm/celeritas/filesystem"
)

type SFTP struct {
	Host     string
	User     string
	Password string
	Port     string
}

func (s *SFTP) Put(filename, folder string) error {
	return nil
}

func (s *SFTP) Get(dst string, items ...string) error {
	return nil
}

func (s *SFTP) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	return entries, nil
}

func (s *SFTP) Delete(toDelete []string) (bool, error) {
	return true, nil
}
