package reader

import (
	"os"

	"github.com/StuartsHome/cmd/config"
)

var _ Reader = &ReaderImpl{}

type Reader interface {
	Read() ([]byte, error)
}

type ReaderImpl struct {
	config *config.Config
}

func New(config *config.Config) *ReaderImpl {
	return &ReaderImpl{
		config: config,
	}
}

func (r *ReaderImpl) Read() ([]byte, error) {
	return os.ReadFile(r.config.InputPath)
}
