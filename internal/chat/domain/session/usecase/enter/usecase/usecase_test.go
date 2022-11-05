package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter/usecase"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter/usecase/mocks"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

var testCases = []testCase{
	{
		"ok",
		tcGive{},
		tcWant{},
	},
	{
		"Repository CreateOrFindUser error",
		tcGive{
			stubs: tcgStubs{
				repository: tcgsRepository{
					createOrFindUser: tcgsRepositoryCreateOrFindUser{
						err: errStubRepositoryCreateOrFindUser,
					},
				},
			},
		},
		tcWant{
			err: errStubRepositoryCreateOrFindUser,
		},
	},
	{
		"Repository CreateOrFindUser and CreateSession erros",
		tcGive{
			stubs: tcgStubs{
				repository: tcgsRepository{
					createOrFindUser: tcgsRepositoryCreateOrFindUser{
						err: errStubRepositoryCreateOrFindUser,
					},
					createSession: tcgsRepositoryCreateSession{
						err: errStubRepositoryCreateSession,
					},
				},
			},
		},
		tcWant{
			err: errStubRepositoryCreateOrFindUser,
		},
	},
	{
		"Repository CreateOrFindUser and CreateSession and MesageBus SubscribeSessionForNewMessages erros",
		tcGive{
			stubs: tcgStubs{
				repository: tcgsRepository{
					createOrFindUser: tcgsRepositoryCreateOrFindUser{
						err: errStubRepositoryCreateOrFindUser,
					},
					createSession: tcgsRepositoryCreateSession{
						err: errStubRepositoryCreateSession,
					},
				},
				messageBus: tcgsMessageBus{
					subscribeSessionForNewMessages: tcgsMessageBusSubscribeSessionForNewMessages{
						err: errStubMessageBusSubscribeSessionForNewMessages,
					},
				},
			},
		},
		tcWant{
			err: errStubRepositoryCreateOrFindUser,
		},
	},
	{
		"Repository CreateSession error",
		tcGive{
			stubs: tcgStubs{
				repository: tcgsRepository{
					createSession: tcgsRepositoryCreateSession{
						err: errStubRepositoryCreateSession,
					},
				},
			},
		},
		tcWant{
			err: errStubRepositoryCreateSession,
		},
	},
	{
		"Repository CreateSession and MessageBus SubscribeSessionForNewMessages errors",
		tcGive{
			stubs: tcgStubs{
				repository: tcgsRepository{
					createSession: tcgsRepositoryCreateSession{
						err: errStubRepositoryCreateSession,
					},
				},
				messageBus: tcgsMessageBus{
					subscribeSessionForNewMessages: tcgsMessageBusSubscribeSessionForNewMessages{
						err: errStubMessageBusSubscribeSessionForNewMessages,
					},
				},
			},
		},
		tcWant{
			err: errStubRepositoryCreateSession,
		},
	},
	{
		"MessageBus SubscribeSessionForNewMessages error",
		tcGive{
			stubs: tcgStubs{
				messageBus: tcgsMessageBus{
					subscribeSessionForNewMessages: tcgsMessageBusSubscribeSessionForNewMessages{
						err: errStubMessageBusSubscribeSessionForNewMessages,
					},
				},
			},
		},
		tcWant{
			err: errStubMessageBusSubscribeSessionForNewMessages,
		},
	},
}

func TestUseCase(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				messageBus = mocks.NewMessageBus(t)
				repository = mocks.NewRepository(t)
				useCase    = usecase.New(messageBus, repository)
			)

			testCase.give.stubs.stubMocks(repository, messageBus)

			gotResult, gotErr := useCase.Do(ctxStub, argsStub)
			if testCase.want.err == nil {
				require.NoError(t, gotErr)
				require.Equal(
					t,
					&enter.Result{
						Messages:  messagesStub,
						SessionID: sessionEntityStub.ID,
					},
					gotResult,
				)
			} else {
				require.ErrorIs(t, gotErr, testCase.want.err)
				require.Nil(t, gotResult)
			}
		})
	}
}

type (
	testCase struct {
		name string
		give tcGive
		want tcWant
	}
)

type (
	tcGive struct {
		stubs tcgStubs
	}

	tcgStubs struct {
		messageBus tcgsMessageBus
		repository tcgsRepository
	}

	tcgsMessageBus struct {
		subscribeSessionForNewMessages tcgsMessageBusSubscribeSessionForNewMessages
	}

	tcgsMessageBusSubscribeSessionForNewMessages struct {
		err error
	}
	tcgsRepository struct {
		createOrFindUser tcgsRepositoryCreateOrFindUser
		createSession    tcgsRepositoryCreateSession
	}

	tcgsRepositoryCreateOrFindUser struct {
		err error
	}

	tcgsRepositoryCreateSession struct {
		err error
	}
)

type (
	tcWant struct {
		err error
	}
)

func (s *tcgStubs) stubMocks(r *mocks.Repository, mb *mocks.MessageBus) {
	s.repository.stubRepositoryMock(r)
	if s.repository.createOrFindUser.err == nil && s.repository.createSession.err == nil {
		s.messageBus.stubMessageBusMock(mb)
	}
}

func (s *tcgsRepository) stubRepositoryMock(r *mocks.Repository) {
	s.stubCreateOrFindUserCall(r)
	if s.createOrFindUser.err == nil {
		s.stubCreateSessionCall(r)
	}
}

func (s *tcgsRepository) stubCreateOrFindUserCall(r *mocks.Repository) {
	var (
		call   = r.On("CreateOrFindUser", ctxStub, argsStub.UserName)
		retErr = s.createOrFindUser.err
		retVal *user.Entity
	)

	if retErr == nil {
		retVal = userEntityStub
	}

	call.Return(retVal, retErr)
}

func (s *tcgsRepository) stubCreateSessionCall(r *mocks.Repository) {
	var (
		call   = r.On("CreateActiveSession", ctxStub, userEntityStub.ID)
		retErr = s.createSession.err
		retVal *session.Entity
	)

	if retErr == nil {
		retVal = sessionEntityStub
	}

	call.Return(retVal, retErr)
}

func (s *tcgsMessageBus) stubMessageBusMock(mb *mocks.MessageBus) {
	s.stubSubscribeSessionForNewMessagesCall(mb)
}

func (s *tcgsMessageBus) stubSubscribeSessionForNewMessagesCall(mb *mocks.MessageBus) {
	var (
		call   = mb.On("SubscribeSessionForNewMessages", ctxStub, sessionEntityStub.ID)
		retErr = s.subscribeSessionForNewMessages.err
		retVal <-chan *message.Entity
	)

	if retErr == nil {
		retVal = messagesStub
	}

	call.Return(retVal, retErr)
}

var (
	ctxStub = context.TODO()

	argsStub = &enter.Args{
		UserName: "username",
	}

	userEntityStub = &user.Entity{
		ID: "stubUserEntityID",
	}

	sessionEntityStub = &session.Entity{
		ID: "stubSessionEntityID",
	}

	messagesStub = make(chan *message.Entity)

	errStubRepositoryCreateOrFindUser               = errors.New("error stub for GatewayRepository CreateOrFindUser")
	errStubRepositoryCreateSession                  = errors.New("error stub for Repository CreateSession")
	errStubMessageBusSubscribeSessionForNewMessages = errors.New("error stub for MessageBus SubscribeSessionForNewMessages")
)
