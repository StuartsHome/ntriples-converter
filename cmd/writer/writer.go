package writer

import (
	"os"

	"github.com/StuartsHome/cmd/config"
)

var _ Writer = &WriterImpl{}

type Writer interface {
	Write(contents []byte, filePath string) error
}

type WriterImpl struct {
	config config.Config
}

func New(config config.Config) *WriterImpl {
	return &WriterImpl{
		config: config,
	}
}

func (w *WriterImpl) Write(contents []byte, filePath string) error {
	return os.WriteFile(filePath, contents, 0666)
}
