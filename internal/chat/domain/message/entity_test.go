package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ID_Int64(t *testing.T) {
	require.Equal(t, int64(1), ID(1).Int64())
	require.NotEqual(t, int64(1), ID(0).Int64())
}
