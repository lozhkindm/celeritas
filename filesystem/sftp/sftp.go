package sftp

import (
	"fmt"
	"github.com/lozhkindm/celeritas/filesystem"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"strings"
)

type SFTP struct {
	Host     string
	User     string
	Password string
	Port     string
}

func (s *SFTP) getCredentials() (*sftp.Client, error) {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User:            s.User,
		Auth:            []ssh.AuthMethod{ssh.Password(s.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, err
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	wd, err := client.Getwd()
	fmt.Println(wd)
	return client, nil
}

func (s *SFTP) Put(filename, folder string) error {
	client, err := s.getCredentials()
	if err != nil {
		return err
	}
	defer func() {
		_ = client.Close()
	}()

	fileToUpload, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = fileToUpload.Close()
	}()

	fileSftp, err := client.Create(path.Base(filename))
	if err != nil {
		return err
	}
	defer func() {
		_ = fileSftp.Close()
	}()

	if _, err := io.Copy(fileSftp, fileToUpload); err != nil {
		return err
	}
	return nil
}

func (s *SFTP) Get(dst string, items ...string) error {
	return nil
}

func (s *SFTP) List(prefix string) ([]filesystem.ListEntry, error) {
	var entries []filesystem.ListEntry
	client, err := s.getCredentials()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = client.Close()
	}()
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

func (s *SFTP) Delete(toDelete []string) (bool, error) {
	return true, nil
}
