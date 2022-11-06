package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ID_String(t *testing.T) {
	require.Equal(t, "1", ID("1").String())
	require.NotEqual(t, "1", ID("").String())
}

func Test_Name_String(t *testing.T) {
	require.Equal(t, "1", Name("1").String())
	require.NotEqual(t, "1", Name("").String())
}
