package logout

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"

	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/logout/gateway/session"
)

type UseCase interface {
	Do(ctx context.Context, request *Request) (*Response, error)
}

type Request struct {
	SessionID session.ID
}

type Response struct {
}

var (
	ErrSessionNotFound = errors.New("session  is not found")
)

func New(
	messageUnsubscriber messageGateway.Unsubscriber,
	sessionDeactivator sessionGateway.Deactivator,
	sessionFinder sessionGateway.Finder,
) UseCase {
	return useCase{
		messageUnsubscriber: messageUnsubscriber,
		sessionDeactivator:  sessionDeactivator,
		sessionFinder:       sessionFinder,
	}
}

type useCase struct {
	messageUnsubscriber messageGateway.Unsubscriber
	sessionDeactivator  sessionGateway.Deactivator
	sessionFinder       sessionGateway.Finder
}

func (uc useCase) Do(ctx context.Context, request *Request) (*Response, error) {
	sessionEntity, err := uc.sessionFinder.Find(ctx, request.SessionID)
	if err != nil {
		return nil, fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return nil, ErrSessionNotFound
	}

	if err := uc.messageUnsubscriber.Unsubscribe(ctx, sessionEntity.ID); err != nil {
		return nil, fmt.Errorf("unsubsribe messages: %w", err)
	}

	if err := uc.sessionDeactivator.Deactivate(ctx, sessionEntity.ID); err != nil {
		return nil, fmt.Errorf("deactivate session: %w", err)
	}

	return &Response{}, nil
}
