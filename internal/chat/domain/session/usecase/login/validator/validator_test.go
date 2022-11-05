package validator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login"
)

func TestArgsValidator(t *testing.T) {
	for _, testCase := range []struct {
		name string
		give login.Args
		want error
	}{
		{
			"ok",
			login.Args{
				UserName: "a",
			},
			nil,
		},
		{
			"empty user name",
			login.Args{
				UserName: "",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			login.Args{
				UserName: " ",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			login.Args{
				UserName: "\t",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			login.Args{
				UserName: "\r\n",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			login.Args{
				UserName: "\n",
			},
			ErrEmptyUserName,
		},
		{
			"empty user name",
			login.Args{
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
