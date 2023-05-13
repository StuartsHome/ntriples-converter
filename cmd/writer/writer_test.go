package writer

import (
	"os"
	"testing"

	"github.com/StuartsHome/cmd/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var outputPath = "./test_data.nt"

func TestWrite_EmptyNamespace(t *testing.T) {
	// Given
	cfg := config.New(config.Namespace(""), "", outputPath)
	writer := New(cfg)

	contents := []byte("sa,b,c\n1,2,3")

	// When
	err := writer.Write(contents)
	require.Nil(t, err)

	// Then
	expected := contents
	got, err := readAndDeleteOutputFile()
	require.Nil(t, err)
	assert.Equal(t, expected, got)
}

func readAndDeleteOutputFile() ([]byte, error) {
	contents, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, err
	}
	err = os.Remove(outputPath)
	if err != nil {
		return nil, err
	}
	return contents, err
}
