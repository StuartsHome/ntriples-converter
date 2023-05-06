package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var path = "fixtures/test_data.csv"

func TestRead(t *testing.T) {
	// Given
	reader := New()

	// When
	got, err := reader.Read(path)

	// Then
	expected := []byte("a,b,c\n1,2,3")
	require.Nil(t, err)
	assert.Equal(t, got, expected)
}
