package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
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

		testCaseGiveStubsRepository struct {
			createOrFindUser createOrFindUserStub
			createSession    createSessionStub
		}

		subscribeForNewMessagesStub struct {
			err error
		}

		testCaseGiveStubsMessageBus struct {
			subscribeForNewMessages subscribeForNewMessagesStub
		}

		testCaseGiveStubs struct {
			repository testCaseGiveStubsRepository
			messageBus testCaseGiveStubsMessageBus
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

		stubMessages = make(chan *message.Entity)

		stubErrRepositoryCreateOrFindUser = errors.New("GatewayRepository CreateOrFindUser error")

		stubErrRepositoryCreateSession = errors.New("Repository CreateSession error")

		stubErrMessageBusSubscribeForNewMessages = errors.New("MessageBus SubscribeForNewMessages error")
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
					repository: testCaseGiveStubsRepository{
						createOrFindUser: createOrFindUserStub{
							err: stubErrRepositoryCreateOrFindUser,
						},
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateOrFindUser,
			},
		},
		{
			"Repository CreateOrFindUser and CreateSession erros",
			testCaseGive{
				stubs: testCaseGiveStubs{
					repository: testCaseGiveStubsRepository{
						createOrFindUser: createOrFindUserStub{
							err: stubErrRepositoryCreateOrFindUser,
						},
						createSession: createSessionStub{
							err: stubErrRepositoryCreateSession,
						},
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateOrFindUser,
			},
		},
		{
			"Repository CreateOrFindUser and CreateSession and MesageBus SubscribeForNewMessages erros",
			testCaseGive{
				stubs: testCaseGiveStubs{
					repository: testCaseGiveStubsRepository{
						createOrFindUser: createOrFindUserStub{
							err: stubErrRepositoryCreateOrFindUser,
						},
						createSession: createSessionStub{
							err: stubErrRepositoryCreateSession,
						},
					},
					messageBus: testCaseGiveStubsMessageBus{
						subscribeForNewMessages: subscribeForNewMessagesStub{
							err: stubErrMessageBusSubscribeForNewMessages,
						},
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
					repository: testCaseGiveStubsRepository{
						createSession: createSessionStub{
							err: stubErrRepositoryCreateSession,
						},
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateSession,
			},
		},
		{
			"Repository CreateSession and MessageBus SubscribeForNewMessages errors",
			testCaseGive{
				stubs: testCaseGiveStubs{
					repository: testCaseGiveStubsRepository{
						createSession: createSessionStub{
							err: stubErrRepositoryCreateSession,
						},
					},
					messageBus: testCaseGiveStubsMessageBus{
						subscribeForNewMessages: subscribeForNewMessagesStub{
							err: stubErrMessageBusSubscribeForNewMessages,
						},
					},
				},
			},
			testCaseWant{
				err: stubErrRepositoryCreateSession,
			},
		},
		{
			"MessageBus SubscribeForNewMessages error",
			testCaseGive{
				stubs: testCaseGiveStubs{
					messageBus: testCaseGiveStubsMessageBus{
						subscribeForNewMessages: subscribeForNewMessagesStub{
							err: stubErrMessageBusSubscribeForNewMessages,
						},
					},
				},
			},
			testCaseWant{
				err: stubErrMessageBusSubscribeForNewMessages,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				messageBus = mocks.NewMessageBus(t)
				repository = mocks.NewRepository(t)
				useCase    = New(messageBus, repository)
			)

			// Setup stubs for Repository CreateOrFindUser
			repositoryCreateOrFindUserCall := repository.
				On("CreateOrFindUser", stubCtx, stubArgs.UserName).
				Once()
			{
				var (
					returnError = testCase.give.stubs.repository.createOrFindUser.err
					returnValue *user.Entity
				)

				if returnError == nil {
					returnValue = stubUserEntity
				}

				repositoryCreateOrFindUserCall.Return(returnValue, returnError)
			}

			// Setup stubs for Repository CreateSession when CreateOrFindUser is expected to be called
			var repositoryCreateSessionCall *mock.Call
			if testCase.give.stubs.repository.createOrFindUser.err == nil {
				repositoryCreateSessionCall = repository.
					On("CreateSession", stubCtx, stubUserEntity.ID).
					NotBefore(repositoryCreateOrFindUserCall)

				var (
					returnError = testCase.give.stubs.repository.createSession.err
					returnValue *session.Entity
				)

				if returnError == nil {
					returnValue = stubSessionEntity
				}

				repositoryCreateSessionCall.Return(returnValue, returnError)
			}

			// Setup stubs for MessageBus SubscribeForNewMessages when Repository
			// CreateOrFindUser and CreateSession are expected to be called.
			if testCase.give.stubs.repository.createOrFindUser.err == nil && testCase.give.stubs.repository.createSession.err == nil {
				messageBusSubscribeForNewMessagesCall := messageBus.
					On("SubscribeForNewMessages", stubCtx, stubSessionEntity.ID).
					NotBefore(repositoryCreateOrFindUserCall).
					NotBefore(repositoryCreateSessionCall)

				var (
					returnError = testCase.give.stubs.messageBus.subscribeForNewMessages.err
					returnValue <-chan *message.Entity
				)

				if returnError == nil {
					returnValue = stubMessages
				}

				messageBusSubscribeForNewMessagesCall.Return(returnValue, returnError)
			}

			gotResult, gotErr := useCase.Do(stubCtx, stubArgs)
			if testCase.want.err == nil {
				require.NoError(t, gotErr)
				require.Equal(t, &enter.Result{Messages: stubMessages, SessionID: stubSessionEntity.ID}, gotResult)
			} else {
				require.ErrorIs(t, gotErr, testCase.want.err)
				require.Nil(t, gotResult)
			}
		})
	}
}
