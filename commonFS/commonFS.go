package FS

import (
	"net/http"
	"os"
)

type CommonFile struct {
	http.File
}

func (f CommonFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

type CommonFS struct {
	http.FileSystem
}

func (s CommonFS) Open(name string) (http.File, error) {
	file, err := s.FileSystem.Open(name)
	return CommonFile{file}, err
}
