package processor

import (
	"testing"

	"github.com/StuartsHome/cmd/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var inputPath = ""
var outputPath = ""

func TestProcess(t *testing.T) {
	// Given
	cfg := config.New(config.Namespace(""), inputPath, outputPath)
	prc := New(cfg)

	contents := []byte("subject,predicate,object\ntestsubject,testpredicate,testobject")
	// When
	got, err := prc.Process(contents)

	// Then
	expected := []byte("<testsubject> <testpredicate> <testobject>\n")
	require.Nil(t, err)
	assert.Equal(t, string(expected), string(got))
}

func TestProcess_HandleTrailingNewLine(t *testing.T) {
	// Given
	cfg := config.New(config.Namespace(""), inputPath, outputPath)
	prc := New(cfg)

	trailingNewLine := "\n"
	contents := []byte("subject,predicate,object\ntestsubject,testpredicate,testobject" + trailingNewLine)
	// When
	got, err := prc.Process(contents)

	// Then
	expected := []byte("<testsubject> <testpredicate> <testobject>\n")
	require.Nil(t, err)
	assert.Equal(t, string(expected), string(got))
}
