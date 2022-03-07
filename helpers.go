package celeritas

import (
	"crypto/rand"
	"os"
)

const (
	randStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321_+"
)

func (c *Celeritas) RandStr(length int) string {
	s, r := make([]rune, length), []rune(randStr)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

func (c *Celeritas) CreateDirIfNotExists(path string) error {
	const mode = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, mode); err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) CreateFileIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(file)
	}
	return nil
}
