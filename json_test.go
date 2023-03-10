package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUtilInt(t *testing.T) {
	value, err := Int("123")

	require.NoError(t, err)
	require.Equal(t, int64(123), value)
}

func TestUtilBool(t *testing.T) {
	value, err := Bool("true")

	require.NoError(t, err)
	require.Equal(t, true, value)
}

func TestUtilFloat(t *testing.T) {
	float, err := Float("1.2345")

	require.NoError(t, err)
	require.Equal(t, 1.2345, float)
}

func TestCopy(t *testing.T) {
	type Kind struct {
		Name string `json:"name"`
		Wow  string `json:"wow"`
	}

	var kind Kind
	err := Copy(&kind, `{"name":"kind","wow":"good"}`)
	require.NoError(t, err)
	require.Equal(t, "kind", kind.Name)
	require.Equal(t, "good", kind.Wow)
}
