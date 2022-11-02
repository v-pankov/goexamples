package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseMessageSend interface {
	DoUseCaseMessageSend(
		ctx context.Context,
		args *UseCaseMessageSendArgs,
	) (
		*UseCaseMessageSendResult,
		error,
	)
}

type UseCaseMessageSendArgs struct {
	AuthorUserID entity.UserID
	RoomID       entity.RoomID
	MessageText  string
}

type UseCaseMessageSendResult struct {
}

func NewUseCaseMessageSend(
	gatewayMessageCreator UseCaseMessageSendGatewayMessageCreator,
	gatewayNewMessageSessionsNotifier UseCaseMessageSendGatewayNewMessageSessionsNotifier,
) UseCaseMessageSend {
	return useCaseMessageSend{
		gatewayMessageCreator:             gatewayMessageCreator,
		gatewayNewMessageSessionsNotifier: gatewayNewMessageSessionsNotifier,
	}
}

type UseCaseMessageSendGatewayMessageCreator interface {
	UseCaseMessageSendGatewayCreateMessage(
		ctx context.Context,
		authorUserID entity.UserID,
		roomID entity.RoomID,
		messageText string,
	) (
		*entity.Message,
		error,
	)
}

type UseCaseMessageSendGatewayNewMessageSessionsNotifier interface {
	UseCaseMessageSendGatewayNotifySessionsAboutNewMessage(
		ctx context.Context,
		nessage *entity.Message,
	) error
}

type useCaseMessageSend struct {
	gatewayMessageCreator             UseCaseMessageSendGatewayMessageCreator
	gatewayNewMessageSessionsNotifier UseCaseMessageSendGatewayNewMessageSessionsNotifier
}

func (uc useCaseMessageSend) DoUseCaseMessageSend(
	ctx context.Context,
	args *UseCaseMessageSendArgs,
) (
	*UseCaseMessageSendResult,
	error,
) {
	messageEntity, err := uc.
		gatewayMessageCreator.
		UseCaseMessageSendGatewayCreateMessage(
			ctx,
			args.AuthorUserID,
			args.RoomID,
			args.MessageText,
		)
	if err != nil {
		return nil, fmt.Errorf("create message: %w", err)
	}

	err = uc.
		gatewayNewMessageSessionsNotifier.
		UseCaseMessageSendGatewayNotifySessionsAboutNewMessage(
			ctx, messageEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new message: %w", err)
	}

	return &UseCaseMessageSendResult{}, nil
}

type UseCaseMessageSendArgsValidator interface {
	ValidateUseCaseMessageSendArgs(ctx context.Context, args *UseCaseMessageSendArgs) error
}

var (
	ErrUseCaseMessageSendArgsAuthorUserDoesNotExist = errors.New("author user does not exist")
	ErrUseCaseMessageSendArgsRoomDoesNotExist       = errors.New("room does not exist")
)

func NewUseCaseMessageSendArgsValidator(
	gatewayUserFinder UseCaseMessageSendArgsValidatorGatewayUserFinder,
	gatewayRoomFinder UseCaseMessageSendArgsValidatorGatewayRoomFinder,
) UseCaseMessageSendArgsValidator {
	return useCaseMessageSendArgsValidator{
		gatewayUserFinder: gatewayUserFinder,
		gatewayRoomFinder: gatewayRoomFinder,
	}
}

type useCaseMessageSendArgsValidator struct {
	gatewayUserFinder UseCaseMessageSendArgsValidatorGatewayUserFinder
	gatewayRoomFinder UseCaseMessageSendArgsValidatorGatewayRoomFinder
}

func (v useCaseMessageSendArgsValidator) ValidateUseCaseMessageSendArgs(
	ctx context.Context, args *UseCaseMessageSendArgs,
) error {
	if err := args.AuthorUserID.Validate(); err != nil {
		return fmt.Errorf("author user id: %w", err)
	}

	if err := args.RoomID.Validate(); err != nil {
		return fmt.Errorf("room id: %w", err)
	}

	authorUserEntity, err := v.
		gatewayUserFinder.
		UseCaseMessageSendArgsValidatorGatewayFindUser(
			ctx, args.AuthorUserID,
		)
	if err != nil {
		return fmt.Errorf("find author user: %w", err)
	}

	if authorUserEntity == nil {
		return ErrUseCaseMessageSendArgsAuthorUserDoesNotExist
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		UseCaseMessageSendArgsValidatorGatewayRoomFinder(
			ctx, args.RoomID,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity == nil {
		return ErrUseCaseMessageSendArgsRoomDoesNotExist
	}

	return nil
}

type UseCaseMessageSendArgsValidatorGatewayUserFinder interface {
	UseCaseMessageSendArgsValidatorGatewayFindUser(
		ctx context.Context, userID entity.UserID,
	) (
		*entity.User,
		error,
	)
}

type UseCaseMessageSendArgsValidatorGatewayRoomFinder interface {
	UseCaseMessageSendArgsValidatorGatewayRoomFinder(
		ctx context.Context, roomID entity.RoomID,
	) (
		*entity.Room,
		error,
	)
}
