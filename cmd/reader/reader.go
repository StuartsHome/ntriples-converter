package reader

import (
	"os"
)

var _ Reader = &ReaderImpl{}

type Reader interface {
	Read(path string) ([]byte, error)
}

type ReaderImpl struct {
}

func New() *ReaderImpl {
	return &ReaderImpl{}
}

func (r *ReaderImpl) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}
