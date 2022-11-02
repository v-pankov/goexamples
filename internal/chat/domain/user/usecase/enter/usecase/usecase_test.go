package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter/usecase/mocks"
)

func TestUseCase(t *testing.T) {
	type (
		gatewayCreateOrFindUserStub struct {
			err error
		}

		gatewayCreateSessionStub struct {
			err error
		}

		testCaseGiveStubs struct {
			gatewayCreateOrFindUser gatewayCreateOrFindUserStub
			gatewayCreateSession    gatewayCreateSessionStub
		}

		testCaseGive struct {
			stubs testCaseGiveStubs
		}

		testCaseWant struct {
			err error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	var (
		stubCtx = context.TODO()

		stubArgs = &enter.Args{
			UserName: "username",
		}

		stubUserEntity = &user.Entity{
			ID: "stubUserEntityID",
		}

		stubSessionEntity = &session.Entity{
			ID: "stubSessionEntityID",
		}

		stubErrGatewayCreateOrFindUser = errors.New("GatewayCreateOrFindUser error")

		stubErrGatewayCreateSession = errors.New("GatewayCreateSession error")
	)

	testCases := []testCase{
		{
			"ok",
			testCaseGive{},
			testCaseWant{},
		},
		{
			"GatewayCreateOrFindUser fails",
			testCaseGive{
				stubs: testCaseGiveStubs{
					gatewayCreateOrFindUser: gatewayCreateOrFindUserStub{
						err: stubErrGatewayCreateOrFindUser,
					},
				},
			},
			testCaseWant{
				err: stubErrGatewayCreateOrFindUser,
			},
		},
		{
			"GatewayCreateSession fails",
			testCaseGive{
				stubs: testCaseGiveStubs{
					gatewayCreateSession: gatewayCreateSessionStub{
						err: stubErrGatewayCreateSession,
					},
				},
			},
			testCaseWant{
				err: stubErrGatewayCreateSession,
			},
		},
		{
			"GatewayCreateOrFindUser and GatewayCreateSession fail",
			testCaseGive{
				stubs: testCaseGiveStubs{
					gatewayCreateOrFindUser: gatewayCreateOrFindUserStub{
						err: stubErrGatewayCreateOrFindUser,
					},
					gatewayCreateSession: gatewayCreateSessionStub{
						err: stubErrGatewayCreateSession,
					},
				},
			},
			testCaseWant{
				err: stubErrGatewayCreateOrFindUser,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Setup GatewayCreateOrFindUser expectations
			var (
				gatewayCreateOrFindUserMock     = mocks.NewGatewayCreateOrFindUser(t)
				gatewayCreateOrFindUserMockCall = gatewayCreateOrFindUserMock.On("Call", stubCtx, stubArgs.UserName)
			)
			if testCase.give.stubs.gatewayCreateOrFindUser.err == nil {
				gatewayCreateOrFindUserMockCall.Return(stubUserEntity, nil)
			} else {
				gatewayCreateOrFindUserMockCall.Return(nil, testCase.give.stubs.gatewayCreateOrFindUser.err)
			}

			// Setup GatewayCreateSession expectations
			var (
				gatewayCreateSessionMock = mocks.NewGatewayCreateSession(t)
			)
			// Expect GatewayCreateSession call when GatewayCreateOrFindUser is expected to succeed.
			if testCase.give.stubs.gatewayCreateOrFindUser.err == nil {
				gatewayCreateSessionMockCall := gatewayCreateSessionMock.On("Call", stubCtx, stubUserEntity.ID)
				if testCase.give.stubs.gatewayCreateSession.err == nil {
					gatewayCreateSessionMockCall.Return(stubSessionEntity, nil)
				} else {
					gatewayCreateSessionMockCall.Return(nil, testCase.give.stubs.gatewayCreateSession.err)
				}
			}

			useCase := New(
				gatewayCreateOrFindUserMock,
				gatewayCreateSessionMock,
			)

			gotResult, gotErr := useCase.Do(stubCtx, stubArgs)
			if testCase.want.err == nil {
				require.NoError(t, gotErr)
				require.Equal(t, &enter.Result{SessionID: stubSessionEntity.ID}, gotResult)
			} else {
				require.ErrorIs(t, gotErr, testCase.want.err)
				require.Nil(t, gotResult)
			}
		})
	}
}
