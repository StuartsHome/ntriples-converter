package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcess(t *testing.T) {
	// Given
	prc := New()
	// When
	got, err := prc.Process(nil)

	// Then
	expected := ""
	require.Nil(t, err)
	assert.Equal(t, got, expected)
}
