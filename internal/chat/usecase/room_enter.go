package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseRoomEnter interface {
	DoUseCaseRoomEnter(
		ctx context.Context,
		args *UseCaseRoomEnterArgs,
	) (
		*UseCaseRoomEnterResult,
		error,
	)
}

type UseCaseRoomEnterArgs struct {
	SessionID entity.SessionID
	RoomID    entity.RoomID
}

type UseCaseRoomEnterResult struct {
}

func NewUseCaseRoomEnter(
	gatewaySessionRoomMessagesSubscriber UseCaseRoomEnterGatewaySessionRoomMessagesSubscriber,
) useCaseRoomEnter {
	return useCaseRoomEnter{
		gatewaySessionRoomMessagesSubscriber: gatewaySessionRoomMessagesSubscriber,
	}
}

type UseCaseRoomEnterGatewaySessionRoomMessagesSubscriber interface {
	UseCaseRoomEnterGatewaySubscribeSessionForRoomMessages(
		ctx context.Context,
		sessionID entity.SessionID,
		roomID entity.RoomID,
	) error
}

type useCaseRoomEnter struct {
	gatewaySessionRoomMessagesSubscriber UseCaseRoomEnterGatewaySessionRoomMessagesSubscriber
}

func (uc useCaseRoomEnter) DoUseCaseEnterRoom(
	ctx context.Context,
	args *UseCaseRoomEnterArgs,
) (
	*UseCaseRoomEnterResult,
	error,
) {
	err := uc.
		gatewaySessionRoomMessagesSubscriber.
		UseCaseRoomEnterGatewaySubscribeSessionForRoomMessages(
			ctx, args.SessionID, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("subscribe session for room messages: %w", err)
	}

	return &UseCaseRoomEnterResult{}, nil
}

type UseCaseRoomEnterArgsValidator interface {
	ValidateUseCaseRoomEnterArgs(ctx context.Context, args *UseCaseRoomEnterArgs) error
}

var (
	ErrUseCaseRoomEnterArgsValidatorSessionDoesNotExist = errors.New("session does not exist")
	ErrUseCaseRoomEnterArgsValidatorRoomDoesNotExist    = errors.New("room does not exist")
)

func NewUseCaseRoomEnterArgsValidator(
	gatewaySessionFinder UseCaseRoomEnterArgsValidatorGatewaySessionFinder,
	gatewayRoomFinder UseCaseRoomEnterArgsValidatorGatewayRoomFinder,
) UseCaseRoomEnterArgsValidator {
	return useCaseRoomEnterArgsValidator{
		gatewaySessionFinder: gatewaySessionFinder,
		gatewayRoomFinder:    gatewayRoomFinder,
	}
}

type UseCaseRoomEnterArgsValidatorGatewaySessionFinder interface {
	UseCaseRoomEnterArgsValidatorGatewayFindSession(
		ctx context.Context, sessionID entity.SessionID,
	) (
		*entity.Session,
		error,
	)
}

type UseCaseRoomEnterArgsValidatorGatewayRoomFinder interface {
	UseCaseRoomEnterArgsValidatorGatewayFindRoom(
		ctx context.Context, roomID entity.RoomID,
	) (
		*entity.Room,
		error,
	)
}

type useCaseRoomEnterArgsValidator struct {
	gatewaySessionFinder UseCaseRoomEnterArgsValidatorGatewaySessionFinder
	gatewayRoomFinder    UseCaseRoomEnterArgsValidatorGatewayRoomFinder
}

func (v useCaseRoomEnterArgsValidator) ValidateUseCaseRoomEnterArgs(
	ctx context.Context, args *UseCaseRoomEnterArgs,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	sessionEntity, err := v.
		gatewaySessionFinder.
		UseCaseRoomEnterArgsValidatorGatewayFindSession(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrUseCaseRoomEnterArgsValidatorSessionDoesNotExist
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		UseCaseRoomEnterArgsValidatorGatewayFindRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity == nil {
		return ErrUseCaseRoomEnterArgsValidatorRoomDoesNotExist
	}

	return nil
}
