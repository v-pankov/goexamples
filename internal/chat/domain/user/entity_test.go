package user

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

func TestValidateName(t *testing.T) {
	validName := Name("somename")
	require.NoError(t, validName.Validate(), "valid name must be validated successfuly")

	emptyName := Name("")
	require.ErrorIs(t, emptyName.Validate(), ErrEmptyName, "empty name must fail validation check")

	spaceName := Name(" \t\n\r\n")
	require.ErrorIs(t, spaceName.Validate(), ErrEmptyName, "name consisting from spaces only must fail validation check")
}
