package inmem

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/vdrpkv/goexamples/internal/pkg/entity"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"
	"github.com/vdrpkv/goexamples/internal/chat/gateway/repository"
)

type InMem struct {
	sync.Mutex

	messageMap map[message.ID]*message.Entity
	sessionMap map[session.ID]*session.Entity
	userMap    map[user.ID]*user.Entity

	nextMessageID int64
}

func New() *InMem {
	return &InMem{
		messageMap: make(map[message.ID]*message.Entity),
		sessionMap: make(map[session.ID]*session.Entity),
		userMap:    make(map[user.ID]*user.Entity),
	}
}

func (inmem *InMem) Gateways() repository.Gateways {
	return repository.Gateways{
		UseCaseMessageSendGateways: repository.UseCaseMessageSendGateways{
			MessageCreator: useCaseMessageSendMessageCreatorGatewayAdapter{inmem},
		},
		UseCaseSessionAuthenticateGateways: repository.UseCaseSessionAuthenticateGateways{
			SessionFinder: useCaseSessionAuthenticateSessionFinderGatewayAdapter{inmem},
		},
		UseCaseUserLoginGateways: repository.UseCaseUserLoginGateways{
			SessionCreator:    useCaseUserLoginSessionCreatorGatewayAdapter{inmem},
			UserCreatorFinder: useCaseUserLoginUserCreatorFinderGatewayAdapter{inmem},
		},
		UseCaseUserLogoutGateways: repository.UseCaseUserLogoutGateways{
			SessionDeactivator: useCaseUserLogoutSessionDeactivatorGateway{inmem},
		},
	}
}

type useCaseMessageSendMessageCreatorGatewayAdapter struct {
	*InMem
}

func (a useCaseMessageSendMessageCreatorGatewayAdapter) Create(
	ctx context.Context, sessionID session.ID, messageText string,
) (*message.Entity, error) {
	a.Lock()
	val, err := a.create(ctx, sessionID, messageText)
	a.Unlock()
	return val, err
}

func (a useCaseMessageSendMessageCreatorGatewayAdapter) create(
	_ context.Context, sessionID session.ID, messageText string,
) (*message.Entity, error) {
	if sessionEntity := a.sessionMap[sessionID]; sessionEntity == nil {
		return nil, errors.New("session does not exist")
	}

	a.nextMessageID++

	messageEntity := &message.Entity{
		Entity:    entity.New(),
		ID:        message.ID(a.nextMessageID),
		SessionID: sessionID,
		Text:      messageText,
	}

	a.messageMap[messageEntity.ID] = messageEntity

	return messageEntity, nil
}

type useCaseSessionAuthenticateSessionFinderGatewayAdapter struct {
	*InMem
}

func (a useCaseSessionAuthenticateSessionFinderGatewayAdapter) Find(
	ctx context.Context, sessionID session.ID,
) (*session.Entity, error) {
	a.Lock()
	val, err := a.find(ctx, sessionID)
	a.Unlock()
	return val, err
}

func (a useCaseSessionAuthenticateSessionFinderGatewayAdapter) find(
	_ context.Context, sessionID session.ID,
) (*session.Entity, error) {
	return a.sessionMap[sessionID], nil
}

type useCaseUserLoginSessionCreatorGatewayAdapter struct {
	*InMem
}

func (a useCaseUserLoginSessionCreatorGatewayAdapter) Create(
	ctx context.Context, userID user.ID,
) (*session.Entity, error) {
	a.Lock()
	val, err := a.create(ctx, userID)
	a.Unlock()
	return val, err
}

func (a useCaseUserLoginSessionCreatorGatewayAdapter) create(
	_ context.Context, userID user.ID,
) (*session.Entity, error) {
	if userEntity := a.userMap[userID]; userEntity == nil {
		return nil, errors.New("user does not exist")
	}

	sessionEntity := &session.Entity{
		Entity: entity.New(),
		ID:     session.ID(uuid.New().String()),
		UserID: userID,
		Active: true,
	}

	a.sessionMap[sessionEntity.ID] = sessionEntity

	return sessionEntity, nil
}

type useCaseUserLoginUserCreatorFinderGatewayAdapter struct {
	*InMem
}

func (a useCaseUserLoginUserCreatorFinderGatewayAdapter) CreateOrFind(
	ctx context.Context, userName user.Name,
) (*user.Entity, error) {
	a.Lock()
	val, err := a.createOrFind(ctx, userName)
	a.Unlock()
	return val, err
}

func (a useCaseUserLoginUserCreatorFinderGatewayAdapter) createOrFind(
	_ context.Context, userName user.Name,
) (*user.Entity, error) {
	for _, userEntity := range a.userMap {
		if userEntity.Name == userName {
			return userEntity, nil
		}
	}

	userEntity := &user.Entity{
		Entity: entity.New(),
		ID:     user.ID(uuid.New().String()),
		Name:   userName,
	}

	a.userMap[userEntity.ID] = userEntity

	return userEntity, nil
}

type useCaseUserLogoutSessionDeactivatorGateway struct {
	*InMem
}

func (a useCaseUserLogoutSessionDeactivatorGateway) Deactivate(
	ctx context.Context, sessionID session.ID,
) error {
	a.Lock()
	err := a.deactivate(ctx, sessionID)
	a.Unlock()
	return err
}

func (a useCaseUserLogoutSessionDeactivatorGateway) deactivate(
	_ context.Context, sessionID session.ID,
) error {
	sessionEntity := a.sessionMap[sessionID]

	if sessionEntity == nil {
		return errors.New("session does not exist")
	}

	sessionEntity.Active = false
	sessionEntity.UpdatedAt = time.Now()
	return nil
}
