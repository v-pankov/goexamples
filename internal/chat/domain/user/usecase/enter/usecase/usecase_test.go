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
		createOrFindUserStub struct {
			err error
		}

		createSessionStub struct {
			err error
		}

		testCaseGiveStubs struct {
			createOrFindUser createOrFindUserStub
			createSession    createSessionStub
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

		stubErrRepositoryCreateOrFindUser = errors.New("GatewayRepository CreateOrFindUser error")

		stubErrRepositoryCreateSession = errors.New("Repository CreateSession error")
	)

	testCases := []testCase{
		{
			"ok",
			testCaseGive{},
			testCaseWant{},
		},
		{
			"Repository CreateOrFindUser error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					createOrFindUser: createOrFindUserStub{
						err: stubErrRepositoryCreateOrFindUser,
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateOrFindUser,
			},
		},
		{
			"Repository CreateSession error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					createSession: createSessionStub{
						err: stubErrRepositoryCreateSession,
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateSession,
			},
		},
		{
			"Repository CreateOrFindUser and CreateSession erros",
			testCaseGive{
				stubs: testCaseGiveStubs{
					createOrFindUser: createOrFindUserStub{
						err: stubErrRepositoryCreateOrFindUser,
					},
					createSession: createSessionStub{
						err: stubErrRepositoryCreateSession,
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateOrFindUser,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				repository = mocks.NewRepository(t)
				useCase    = New(repository)
			)

			// Setup stubs for CreateOrFindUser
			repositoryCreateOrFindUserCall := repository.
				On("CreateOrFindUser", stubCtx, stubArgs.UserName).
				Once()
			{
				var (
					returnError = testCase.give.stubs.createOrFindUser.err
					returnValue *user.Entity
				)

				if returnError == nil {
					returnValue = stubUserEntity
				}

				repositoryCreateOrFindUserCall.Return(returnValue, returnError)
			}

			// Setup stubs for CreateSession when CreateOrFindUser is expected to be called
			if testCase.give.stubs.createOrFindUser.err == nil {
				repositoryCreateSessionCall := repository.
					On("CreateSession", stubCtx, stubUserEntity.ID).
					NotBefore(repositoryCreateOrFindUserCall)

				var (
					returnError = testCase.give.stubs.createSession.err
					returnValue *session.Entity
				)

				if returnError == nil {
					returnValue = stubSessionEntity
				}

				repositoryCreateSessionCall.Return(returnValue, returnError)
			}

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
