package reader

import (
	"os"

	"github.com/StuartsHome/cmd/config"
)

var _ Reader = &ReaderImpl{}

type Reader interface {
	Read(path string) ([]byte, error)
}

type ReaderImpl struct {
	config config.Config
}

func New(config config.Config) *ReaderImpl {
	return &ReaderImpl{
		config: config,
	}
}

func (r *ReaderImpl) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}
