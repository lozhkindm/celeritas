package webdav

import (
	"github.com/lozhkindm/celeritas/filesystem"
	"os"
	"path"
	"strings"

	"github.com/studio-b12/gowebdav"
)

type WebDAV struct {
	Host     string
	User     string
	Password string
}

func (w *WebDAV) Put(filename, folder string) error {
	client := gowebdav.NewClient(w.Host, w.User, w.Password)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	if err := client.WriteStream(path.Join(folder, path.Base(filename)), file, 0644); err != nil {
		return err
	}
	return nil
}

func (w *WebDAV) Get(dst string, items ...string) error {
	return nil
}

func (w *WebDAV) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	client := gowebdav.NewClient(w.Host, w.User, w.Password)
	files, err := client.ReadDir(prefix)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			entries = append(entries, filesystem.ListEntry{
				LastModified: file.ModTime(),
				Key:          file.Name(),
				Size:         float64(file.Size()) / 1024 / 1024, // MB
				IsDir:        file.IsDir(),
			})
		}
	}
	return entries, nil
}

func (w *WebDAV) Delete(toDelete []string) (bool, error) {
	return true, nil
}
