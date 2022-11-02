package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseRoomExit interface {
	DoUseCaseRoomExit(
		ctx context.Context,
		args *UseCaseRoomExitArgs,
	) (
		*UseCaseRoomExitResult,
		error,
	)
}

type UseCaseRoomExitArgs struct {
	SessionID entity.SessionID
	RoomID    entity.RoomID
}

type UseCaseRoomExitResult struct {
}

func NewUseCaseRoomExit(
	gatewaySessionRoomMessagesUnsubscriber UseCaseRoomExitGatewaySessionRoomMessagesUnsubscriber,
) useCaseRoomExit {
	return useCaseRoomExit{
		gatewaySessionRoomMessagesUnsubscriber: gatewaySessionRoomMessagesUnsubscriber,
	}
}

type UseCaseRoomExitGatewaySessionRoomMessagesUnsubscriber interface {
	UseCaseRoomExitGatewayUnsubscribeSessionForRoomMessages(
		ctx context.Context,
		sessionID entity.SessionID,
		roomID entity.RoomID,
	) error
}

type useCaseRoomExit struct {
	gatewaySessionRoomMessagesUnsubscriber UseCaseRoomExitGatewaySessionRoomMessagesUnsubscriber
}

func (uc useCaseRoomExit) DoUseCaseExitRoom(
	ctx context.Context,
	args *UseCaseRoomExitArgs,
) (
	*UseCaseRoomExitResult,
	error,
) {
	err := uc.
		gatewaySessionRoomMessagesUnsubscriber.
		UseCaseRoomExitGatewayUnsubscribeSessionForRoomMessages(
			ctx, args.SessionID, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("unsubscribe session for room messages: %w", err)
	}

	return &UseCaseRoomExitResult{}, nil
}

type UseCaseRoomExitArgsValidator interface {
	ValidateUseCaseRoomExitArgs(ctx context.Context, args *UseCaseRoomExitArgs) error
}

var (
	ErrUseCaseRoomExitArgsValidatorSessionDoesNotExist = errors.New("session does not exist")
	ErrUseCaseRoomExitArgsValidatorRoomDoesNotExist    = errors.New("room does not exist")
)

func NewUseCaseRoomExitArgsValidator(
	gatewaySessionFinder UseCaseRoomExitArgsValidatorGatewaySessionFinder,
	gatewayRoomFinder UseCaseRoomExitArgsValidatorGatewayRoomFinder,
) UseCaseRoomExitArgsValidator {
	return useCaseRoomExitArgsValidator{
		gatewaySessionFinder: gatewaySessionFinder,
		gatewayRoomFinder:    gatewayRoomFinder,
	}
}

type UseCaseRoomExitArgsValidatorGatewaySessionFinder interface {
	UseCaseRoomExitArgsValidatorGatewayFindSession(
		ctx context.Context, sessionID entity.SessionID,
	) (
		*entity.Session,
		error,
	)
}

type UseCaseRoomExitArgsValidatorGatewayRoomFinder interface {
	UseCaseRoomExitArgsValidatorGatewayFindRoom(
		ctx context.Context, roomID entity.RoomID,
	) (
		*entity.Room,
		error,
	)
}

type useCaseRoomExitArgsValidator struct {
	gatewaySessionFinder UseCaseRoomExitArgsValidatorGatewaySessionFinder
	gatewayRoomFinder    UseCaseRoomExitArgsValidatorGatewayRoomFinder
}

func (v useCaseRoomExitArgsValidator) ValidateUseCaseRoomExitArgs(
	ctx context.Context, args *UseCaseRoomExitArgs,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	sessionEntity, err := v.
		gatewaySessionFinder.
		UseCaseRoomExitArgsValidatorGatewayFindSession(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrUseCaseRoomExitArgsValidatorSessionDoesNotExist
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		UseCaseRoomExitArgsValidatorGatewayFindRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity == nil {
		return ErrUseCaseRoomExitArgsValidatorRoomDoesNotExist
	}

	return nil
}
