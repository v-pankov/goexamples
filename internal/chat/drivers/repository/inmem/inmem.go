package inmem

import (
	"context"
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"

	messageSendUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/usecase"
	messageSendValidator "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/validator"

	sessionLoginUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login/usecase"

	sessionLogoutUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/usecase"
	sessionLogoutValidator "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/validator"
)

type Repository interface {
	SessionLoginUsecaseRepository() sessionLoginUsecase.Repository
	SessionLogoutUsecaseRepository() sessionLogoutUsecase.Repository

	SessionLogoutValidatorRepository() sessionLogoutValidator.Repository

	MessageSendUsecaseRepository() messageSendUsecase.Repository
	MessageSendValidatorRepository() messageSendValidator.Repository
}

func New() Repository {
	return &repo{
		userMap:    make(map[user.ID]*user.Entity),
		sessionMap: make(map[session.ID]*session.Entity),
		messageMap: make(map[message.ID]*message.Entity),
	}
}

type repo struct {
	sync.Mutex
	userMap    map[user.ID]*user.Entity
	sessionMap map[session.ID]*session.Entity
	messageMap map[message.ID]*message.Entity
}

func (r *repo) SessionLoginUsecaseRepository() sessionLoginUsecase.Repository {
	return sessionLoginUsecaseRepositoryAdapter{r}
}

func (r *repo) SessionLogoutUsecaseRepository() sessionLogoutUsecase.Repository {
	return sessionLogoutUsecaseRepositoryAdapter{r}
}

func (r *repo) SessionLogoutValidatorRepository() sessionLogoutValidator.Repository {
	return sessionLogoutValidatorRepositoryAdapter{r}
}

func (r *repo) MessageSendUsecaseRepository() messageSendUsecase.Repository {
	return messageSendUsecaseRepositoryAdapter{r}
}

func (r *repo) MessageSendValidatorRepository() messageSendValidator.Repository {
	return messageSendValidatorRepositoryAdapter{r}
}

type sessionLoginUsecaseRepositoryAdapter struct {
	*repo
}

func (a sessionLoginUsecaseRepositoryAdapter) CreateOrFindUser(ctx context.Context, userName user.Name) (*user.Entity, error) {
	return nil, nil
}

func (a sessionLoginUsecaseRepositoryAdapter) CreateActiveSession(ctx context.Context, userID user.ID) (*session.Entity, error) {
	return nil, nil
}

type sessionLogoutUsecaseRepositoryAdapter struct {
	*repo
}

func (a sessionLogoutUsecaseRepositoryAdapter) DeactivateSession(ctx context.Context, sessionID session.ID) error {
	return nil
}

type sessionLogoutValidatorRepositoryAdapter struct {
	*repo
}

func (a sessionLogoutValidatorRepositoryAdapter) FindSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	return nil, nil
}

type messageSendUsecaseRepositoryAdapter struct {
	*repo
}

func (a messageSendUsecaseRepositoryAdapter) CreateMessage(
	ctx context.Context,
	authorUserSessionID session.ID,
	messageText string,
) (
	*message.Entity,
	error,
) {
	return nil, nil
}

type messageSendValidatorRepositoryAdapter struct {
	*repo
}

func (a messageSendValidatorRepositoryAdapter) FindSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	return nil, nil
}
