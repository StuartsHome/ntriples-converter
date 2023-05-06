package processor

type Processor interface {
	Process(contents []byte) ([]byte, error)
}

type ProcessorImpl struct {
}

var _ Processor = &ProcessorImpl{}

func New() *ProcessorImpl {
	return &ProcessorImpl{}
}

func (p *ProcessorImpl) Process(contents []byte) ([]byte, error) {
	return contents, nil
}
