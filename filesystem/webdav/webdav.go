package webdav

import (
	"github.com/lozhkindm/celeritas/filesystem"
)

type WebDAV struct {
	Host     string
	User     string
	Password string
}

func (w *WebDAV) Put(filename, folder string) error {
	return nil
}

func (w *WebDAV) Get(dst string, items ...string) error {
	return nil
}

func (w *WebDAV) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	return entries, nil
}

func (w *WebDAV) Delete(toDelete []string) (bool, error) {
	return true, nil
}
