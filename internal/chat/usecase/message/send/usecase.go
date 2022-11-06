package send

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity/session"

	messageGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/message"
	sessionGateway "github.com/vdrpkv/goexamples/internal/chat/usecase/message/send/gateway/session"
)

type Request struct {
	AuthorUserSessionID session.ID
	MessageText         string
}

type Response struct {
}

var (
	ErrAuthorUserSessionNotFound = errors.New("author user session is not found")
)

type UseCase interface {
	Do(ctx context.Context, request *Request) (*Response, error)
}

func New(
	messageBroadcaster messageGateway.Broadcaster,
	messageCreator messageGateway.Creator,
	sessionFinder sessionGateway.Finder,
) UseCase {
	return useCase{
		messageBroadcaster: messageBroadcaster,
		messageCreator:     messageCreator,
		sessionFinder:      sessionFinder,
	}
}

type useCase struct {
	messageBroadcaster messageGateway.Broadcaster
	messageCreator     messageGateway.Creator
	sessionFinder      sessionGateway.Finder
}

func (uc useCase) Do(ctx context.Context, request *Request) (*Response, error) {
	authorUserSessionEntity, err := uc.sessionFinder.Find(ctx, request.AuthorUserSessionID)
	if err != nil {
		return nil, fmt.Errorf("find author user session: %w", err)
	}

	if authorUserSessionEntity == nil {
		return nil, ErrAuthorUserSessionNotFound
	}

	messageEntity, err := uc.messageCreator.Create(
		ctx,
		request.AuthorUserSessionID,
		request.MessageText,
	)

	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	err = uc.messageBroadcaster.Broadcast(ctx, messageEntity)

	if err != nil {
		return nil, fmt.Errorf("broadcast message to all sessions: %w", err)
	}

	return &Response{}, nil
}
