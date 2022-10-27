package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseRoomDelete interface {
	DoUseCaseRoomDelete(
		ctx context.Context,
		args *UseCaseRoomDeleteArgs,
	) (
		*UseCaseRoomDeleteResult,
		error,
	)
}

type UseCaseRoomDeleteArgs struct {
	RoomID entity.RoomID
}

type UseCaseRoomDeleteResult struct {
}

func NewUseCaseRoomDelete(
	gatewaySessionsRoomMessagesUnsubscriber UseCaseRoomDeleteGatewaySessionsRoomMessagesUnsubscriber,
	gatewayRoomDeleter UseCaseRoomDeleteGatewayRoomDeleter,
	gatewaySessionsRoomRemovalNotifier UseCaseRoomDeleteGatewaySessionsRoomRemovalNotifier,
) UseCaseRoomDelete {
	return useCaseRoomDelete{
		gatewaySessionsRoomMessagesUnsubscriber: gatewaySessionsRoomMessagesUnsubscriber,
		gatewayRoomDeleter:                      gatewayRoomDeleter,
		gatewaySessionsRoomRemovalNotifier:      gatewaySessionsRoomRemovalNotifier,
	}
}

type UseCaseRoomDeleteGatewaySessionsRoomMessagesUnsubscriber interface {
	UseCaseRoomDeleteGatewayUnsubscribeSessionsFromRoomMessages(
		ctx context.Context, roomID entity.RoomID,
	) error
}

type UseCaseRoomDeleteGatewayRoomDeleter interface {
	UseCaseRoomDeleteGatewayDeleteRoom(
		ctx context.Context, roomID entity.RoomID,
	) error
}

type UseCaseRoomDeleteGatewaySessionsRoomRemovalNotifier interface {
	UseCaseRoomDeleteGatewayNotifySessionsAboutRemovedRoom(
		ctx context.Context, roomID entity.RoomID,
	) error
}

type useCaseRoomDelete struct {
	gatewaySessionsRoomMessagesUnsubscriber UseCaseRoomDeleteGatewaySessionsRoomMessagesUnsubscriber
	gatewayRoomDeleter                      UseCaseRoomDeleteGatewayRoomDeleter
	gatewaySessionsRoomRemovalNotifier      UseCaseRoomDeleteGatewaySessionsRoomRemovalNotifier
}

func (uc useCaseRoomDelete) DoUseCaseRoomDelete(
	ctx context.Context,
	args *UseCaseRoomDeleteArgs,
) (
	*UseCaseRoomDeleteResult,
	error,
) {
	err := uc.
		gatewaySessionsRoomMessagesUnsubscriber.
		UseCaseRoomDeleteGatewayUnsubscribeSessionsFromRoomMessages(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("unsubsribe sessions from room messages: %w", err)
	}

	err = uc.
		gatewayRoomDeleter.
		UseCaseRoomDeleteGatewayDeleteRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("delete room: %w", err)
	}

	err = uc.
		gatewaySessionsRoomRemovalNotifier.
		UseCaseRoomDeleteGatewayNotifySessionsAboutRemovedRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about removed room: %w", err)
	}

	return nil, nil
}

type UseCaseRoomDeleteArgsValidator interface {
	ValidateUseCaseRoomDeleteArgs(ctx context.Context, args *UseCaseRoomDeleteArgs) error
}

var (
	ErrUseCaseRoomDeleteArgsRoomDoesNotExist = errors.New("room does not exist")
)

func NewUseCaseRoomDeleteArgsValidator(
	gatewayRoomFinder UseCaseRoomDeleteArgsValidatorGatewayRoomFinder,
) UseCaseRoomDeleteArgsValidator {
	return useCaseRoomDeleteArgsValidator{
		gatewayRoomFinder: gatewayRoomFinder,
	}
}

type useCaseRoomDeleteArgsValidator struct {
	gatewayRoomFinder UseCaseRoomDeleteArgsValidatorGatewayRoomFinder
}

func (v useCaseRoomDeleteArgsValidator) ValidateUseCaseRoomDeleteArgs(
	ctx context.Context, args *UseCaseRoomDeleteArgs,
) error {
	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		UseCaseRoomDeleteArgsValidatorGatewayFindRoom(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity != nil {
		return ErrUseCaseRoomDeleteArgsRoomDoesNotExist
	}

	return nil
}

type UseCaseRoomDeleteArgsValidatorGatewayRoomFinder interface {
	UseCaseRoomDeleteArgsValidatorGatewayFindRoom(
		ctx context.Context, roomID entity.RoomID,
	) (
		*entity.Room,
		error,
	)
}
