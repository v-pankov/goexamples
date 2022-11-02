package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ID_String(t *testing.T) {
	require.Equal(t, "1", ID("1").String())
	require.NotEqual(t, "1", ID("").String())
}
