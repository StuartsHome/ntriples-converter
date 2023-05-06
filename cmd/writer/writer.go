package writer

import (
	"os"
)

var _ Writer = &WriterImpl{}

type Writer interface {
	Write(contents []byte, filePath string) error
}

type WriterImpl struct {
}

func New() *WriterImpl {
	return &WriterImpl{}
}

func (w *WriterImpl) Write(contents []byte, filePath string) error {
	return os.WriteFile(filePath, contents, 0666)
}
