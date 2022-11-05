package inmem

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
	"github.com/vdrpkv/goexamples/internal/pkg/entity"

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
	userMap       map[user.ID]*user.Entity
	sessionMap    map[session.ID]*session.Entity
	messageMap    map[message.ID]*message.Entity
	nextMessageID int64
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
	a.Lock()
	val, err := a.createOrFindUser(ctx, userName)
	a.Unlock()
	return val, err
}

func (a sessionLoginUsecaseRepositoryAdapter) CreateActiveSession(ctx context.Context, userID user.ID) (*session.Entity, error) {
	a.Lock()
	val, err := a.createActiveSession(ctx, userID)
	a.Unlock()
	return val, err
}

func (a sessionLoginUsecaseRepositoryAdapter) createOrFindUser(_ context.Context, userName user.Name) (*user.Entity, error) {
	for _, userEntity := range a.repo.userMap {
		if userEntity.Name == userName {
			return userEntity, nil
		}
	}
	return nil, nil
}

func (a sessionLoginUsecaseRepositoryAdapter) createActiveSession(_ context.Context, userID user.ID) (*session.Entity, error) {
	var (
		timeNow       = time.Now()
		sessionEntity = &session.Entity{
			Entity: entity.Entity{
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			ID:     session.ID(uuid.New().String()),
			UserID: userID,
			Active: true,
		}
	)
	a.sessionMap[sessionEntity.ID] = sessionEntity
	return sessionEntity, nil
}

type sessionLogoutUsecaseRepositoryAdapter struct {
	*repo
}

func (a sessionLogoutUsecaseRepositoryAdapter) DeactivateSession(ctx context.Context, sessionID session.ID) error {
	a.Lock()
	err := a.deactivateSession(ctx, sessionID)
	a.Unlock()
	return err
}

func (a sessionLogoutUsecaseRepositoryAdapter) deactivateSession(ctx context.Context, sessionID session.ID) error {
	sessionEntity := a.sessionMap[sessionID]
	if sessionEntity == nil {
		return fmt.Errorf("session is not found")
	}

	sessionEntity.Active = false
	return nil
}

type sessionLogoutValidatorRepositoryAdapter struct {
	*repo
}

func (a sessionLogoutValidatorRepositoryAdapter) FindSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	a.Lock()
	val, err := a.findSession(ctx, sessionID)
	a.Unlock()
	return val, err
}

func (a sessionLogoutValidatorRepositoryAdapter) findSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	return a.sessionMap[sessionID], nil
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
	a.Lock()
	val, err := a.createMessage(ctx, authorUserSessionID, messageText)
	a.Unlock()
	return val, err
}

func (a messageSendUsecaseRepositoryAdapter) createMessage(
	ctx context.Context,
	authorUserSessionID session.ID,
	messageText string,
) (
	*message.Entity,
	error,
) {
	a.nextMessageID++
	var (
		timeNow       = time.Now()
		messageEntity = &message.Entity{
			Entity: entity.Entity{
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			ID:        message.ID(a.nextMessageID),
			SessionID: authorUserSessionID,
			Text:      messageText,
		}
	)
	a.messageMap[messageEntity.ID] = messageEntity
	return messageEntity, nil
}

type messageSendValidatorRepositoryAdapter struct {
	*repo
}

func (a messageSendValidatorRepositoryAdapter) FindSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	a.Lock()
	val, err := a.findSession(ctx, sessionID)
	a.Unlock()
	return val, err
}

func (a messageSendValidatorRepositoryAdapter) findSession(ctx context.Context, sessionID session.ID) (*session.Entity, error) {
	return a.sessionMap[sessionID], nil
}
