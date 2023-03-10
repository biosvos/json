package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestByteListToMap(t *testing.T) {
	_, err := byteList(exampleJson).toMap()
	require.NoError(t, err)
}

func TestByteListToSlice(t *testing.T) {
	_, err := byteList(`["GML", "XML"]`).toSlice()
	require.NoError(t, err)
}
