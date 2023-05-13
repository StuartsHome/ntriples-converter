package writer

import (
	"os"

	"github.com/StuartsHome/cmd/config"
)

var _ Writer = &WriterImpl{}

type Writer interface {
	Write(contents []byte) error
}

type WriterImpl struct {
	config config.Config
}

func New(config *config.Config) *WriterImpl {
	return &WriterImpl{
		config: *config,
	}
}

func (w *WriterImpl) Write(contents []byte) error {
	return os.WriteFile(w.config.OutputPath, contents, 0666)
}
