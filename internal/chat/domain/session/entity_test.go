package session

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateID(t *testing.T) {
	validID := ID("123")
	require.NoError(t, validID.Validate(), "valid ID must be validated successfuly")

	emptyID := ID("")
	require.ErrorIs(t, emptyID.Validate(), ErrEmptyID, "empty ID must fail validation check")
}
