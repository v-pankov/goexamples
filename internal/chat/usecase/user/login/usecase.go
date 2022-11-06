package login

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/vdrpkv/goexamples/internal/chat/entity/message"
	"github.com/vdrpkv/goexamples/internal/chat/entity/session"
	"github.com/vdrpkv/goexamples/internal/chat/entity/user"

	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/session"
	userGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/user/login/gateway/user"
)

type Request struct {
	UserName user.Name
}

type Response struct {
	Messages  <-chan *message.Entity
	SessionID session.ID
}

var (
	ErrEmptyUserName = errors.New("user name is empty")
)

type UseCase interface {
	Do(ctx context.Context, request *Request) (*Response, error)
}

func New(
	messageSubscriber messageGateway.Subscriber,
	sessionCreator sessionGateway.Creator,
	userCreatorFinder userGateway.CreatorFinder,
) UseCase {
	return useCase{
		messageSubscriber: messageSubscriber,
		sessionCreator:    sessionCreator,
		userCreatorFinder: userCreatorFinder,
	}
}

type useCase struct {
	messageSubscriber messageGateway.Subscriber
	sessionCreator    sessionGateway.Creator
	userCreatorFinder userGateway.CreatorFinder
}

func (uc useCase) Do(ctx context.Context, request *Request) (*Response, error) {
	if len(strings.TrimSpace(request.UserName.String())) == 0 {
		return nil, ErrEmptyUserName
	}

	userEntity, err := uc.userCreatorFinder.CreateOrFind(ctx, request.UserName)
	if err != nil {
		return nil, fmt.Errorf("create or find user [%s]: %w", request.UserName, err)
	}

	sessionEntity, err := uc.sessionCreator.Create(ctx, userEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}

	messages, err := uc.messageSubscriber.Subscribe(ctx, sessionEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("subscribe messages: %w", err)
	}

	return &Response{
		Messages:  messages,
		SessionID: sessionEntity.ID,
	}, nil
}
