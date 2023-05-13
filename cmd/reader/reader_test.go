package reader

import (
	"testing"

	"github.com/StuartsHome/cmd/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var inputPath = "fixtures/test_data.csv"
var outputPath = ""

func TestRead(t *testing.T) {
	// Given
	cfg := config.New(config.Namespace(""), inputPath, outputPath)
	reader := New(cfg)

	// When
	got, err := reader.Read()

	// Then
	expected := []byte("subject,predicate,object\ntestsubject,testpredicate,testobject")
	require.Nil(t, err)
	assert.Equal(t, got, expected)
}
