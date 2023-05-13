package processor

import (
	"fmt"
	"strings"

	"github.com/StuartsHome/cmd/config"
)

type Processor interface {
	Process(contents []byte) ([]byte, error)
}

type ProcessorImpl struct {
	config *config.Config
}

var _ Processor = &ProcessorImpl{}

func New(config *config.Config) *ProcessorImpl {
	return &ProcessorImpl{
		config: config,
	}
}

func (p *ProcessorImpl) Process(contents []byte) ([]byte, error) {
	strContents := string(contents)
	strContentsSlice := strings.Split(strContents, "\n")

	strContentsSlice = strContentsSlice[1:]

	var output string
	for _, row := range strContentsSlice {
		rr := strings.Split(row, ",")
		if len(rr) == 1 && rr[0] == "" {
			continue
		}
		var newRow string
		for _, col := range rr {
			newRow = newRow + fmt.Sprintf(" <%s>", col)
		}

		output = output + strings.Trim(newRow, " ") + "\n"
	}
	return []byte(output), nil
}
