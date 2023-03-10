package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBytesToMap(t *testing.T) {
	_, err := toMap([]byte(exampleJson))
	require.NoError(t, err)
}

func TestBytesToSlice(t *testing.T) {
	_, err := toSlice([]byte(`["GML", "XML"]`))
	require.NoError(t, err)
}
