package ntriples

import (
	"fmt"

	"github.com/StuartsHome/cmd/config"
	"github.com/StuartsHome/cmd/processor"
	"github.com/StuartsHome/cmd/reader"
	"github.com/StuartsHome/cmd/writer"
)

var _ Ntriples = &NtriplesImpl{}

type Ntriples interface {
	Construct(filepath string) error
}

type NtriplesImpl struct {
	reader    reader.Reader
	processor processor.Processor
	writer    writer.Writer
	config.Config
}

func New(config config.Config) *NtriplesImpl {
	return &NtriplesImpl{
		reader:    reader.New(config),
		processor: processor.New(config),
		writer:    writer.New(config),
	}
}

var readerErr = func(err error) error { return fmt.Errorf("unable to read file: %w", err) }
var processorErr = func(err error) error { return fmt.Errorf("unable to process contents: %w", err) }
var writerErr = func(err error) error { return fmt.Errorf("unable to write processed contents: %w", err) }

func (n *NtriplesImpl) Construct(filepath string) error {
	// Read.
	contents, err := n.reader.Read(filepath)
	if err != nil {
		return readerErr(err)
	}

	// Process.
	process, err := n.processor.Process(contents)
	if err != nil {
		return processorErr(err)
	}

	// Write.
	if err := n.writer.Write(process, n.OutputPath); err != nil {
		return writerErr(err)
	}
	return nil
}
