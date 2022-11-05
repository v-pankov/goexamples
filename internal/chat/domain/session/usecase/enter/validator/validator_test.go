package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter"
)

func TestArgsValidator(t *testing.T) {
	for _, testCase := range []struct {
		name string
		give enter.Args
		want error
	}{
		{
			"ok",
			enter.Args{
				UserName: "a",
			},
			nil,
		},
		{
			"empty user name",
			enter.Args{
				UserName: "",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			enter.Args{
				UserName: " ",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			enter.Args{
				UserName: "\t",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			enter.Args{
				UserName: "\r\n",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			enter.Args{
				UserName: "\n",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			enter.Args{
				UserName: " \t\r\n\n",
			},
			ErrEmptyUserName,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				ctx                 = context.TODO()
				argsValidator       = NewArgsValidator()
				argsValidationError = argsValidator.ValidateArgs(ctx, &testCase.give)
			)

			if testCase.want == nil {
				require.NoError(t, argsValidationError)
			} else {
				require.ErrorIs(t, argsValidationError, testCase.want)
			}
		})
	}
}
