package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseRoomCreate interface {
	DoUseCaseRoomCreate(
		ctx context.Context,
		args *UseCaseRoomCreateArgs,
	) (
		*UseCaseRoomCreateResult,
		error,
	)
}

type UseCaseRoomCreateArgs struct {
	CreatorUserID entity.UserID
	RoomName      entity.RoomName
}

type UseCaseRoomCreateResult struct {
	Room *entity.Room
}

func NewUseCase(
	gatewayRoomCreator UseCaseRoomCreateGatewayRoomCreator,
	gatewayNewRoomSessionsNotifier UseCaseRoomCreateGatewayNewRoomSessionsNotifier,
) useCaseRoomCreate {
	return useCaseRoomCreate{
		gatewayRoomCreator:             gatewayRoomCreator,
		gatewayNewRoomSessionsNotifier: gatewayNewRoomSessionsNotifier,
	}
}

type UseCaseRoomCreateGatewayRoomCreator interface {
	UseCaseRoomCreateGatewayCreateRoom(
		ctx context.Context,
		creatorUserID entity.UserID,
		roomName entity.RoomName,
	) (
		*entity.Room,
		error,
	)
}

type UseCaseRoomCreateGatewayNewRoomSessionsNotifier interface {
	UseCaseRoomCreateGatewayNotifySessionsAboutNewRoom(
		ctx context.Context,
		room *entity.Room,
	) error
}

type useCaseRoomCreate struct {
	gatewayRoomCreator             UseCaseRoomCreateGatewayRoomCreator
	gatewayNewRoomSessionsNotifier UseCaseRoomCreateGatewayNewRoomSessionsNotifier
}

func (uc useCaseRoomCreate) DoUseCaseCreateRoom(
	ctx context.Context,
	args *UseCaseRoomCreateArgs,
) (
	*UseCaseRoomCreateResult,
	error,
) {
	roomEntity, err := uc.
		gatewayRoomCreator.
		UseCaseRoomCreateGatewayCreateRoom(
			ctx, args.CreatorUserID, args.RoomName,
		)
	if err != nil {
		return nil, fmt.Errorf("create room: %w", err)
	}

	err = uc.
		gatewayNewRoomSessionsNotifier.
		UseCaseRoomCreateGatewayNotifySessionsAboutNewRoom(
			ctx, roomEntity,
		)
	if err != nil {
		return nil, fmt.Errorf("notify sessions about new room: %w", err)
	}

	return &UseCaseRoomCreateResult{
		Room: roomEntity,
	}, nil
}

type UseCaseRoomCreateArgsValidator interface {
	ValidateUseCaseRoomCreateArgs(ctx context.Context, args *UseCaseRoomCreateArgs) error
}

var (
	ErrUseCaseRoomCreateArgsCreatorUserDoesNotExist = errors.New("creator user does not exist")
	ErrUseCaseRoomCreateArgsRoomNameIsNotUnique     = errors.New("room name is not unique")
)

func NewUseCaseRoomCreateArgsValidator(
	gatewayUserFinder UseCaseRoomCreateArgsValidatorGatewayUserFinder,
	gatewayRoomFinder UseCaseRoomCreateArgsValidatorGatewayRoomFinder,
) UseCaseRoomCreateArgsValidator {
	return useCaseRoomCreateArgsValidator{
		gatewayUserFinder: gatewayUserFinder,
		gatewayRoomFinder: gatewayRoomFinder,
	}
}

type useCaseRoomCreateArgsValidator struct {
	gatewayUserFinder UseCaseRoomCreateArgsValidatorGatewayUserFinder
	gatewayRoomFinder UseCaseRoomCreateArgsValidatorGatewayRoomFinder
}

func (v useCaseRoomCreateArgsValidator) ValidateUseCaseRoomCreateArgs(
	ctx context.Context, args *UseCaseRoomCreateArgs,
) error {
	if err := args.CreatorUserID.Validate(); err != nil {
		return fmt.Errorf("creator user id: %w", err)
	}

	if err := args.RoomName.Validate(); err != nil {
		return fmt.Errorf("room name: %w", err)
	}

	creatorUserEntity, err := v.
		gatewayUserFinder.
		UseCaseRoomCreateArgsValidatorGatewayFindUser(
			ctx, args.CreatorUserID,
		)
	if err != nil {
		return fmt.Errorf("find creator user: %w", err)
	}

	if creatorUserEntity == nil {
		return ErrUseCaseRoomCreateArgsCreatorUserDoesNotExist
	}

	roomEntity, err := v.
		gatewayRoomFinder.
		UseCaseRoomCreateArgsValidatorGatewayFindRoom(
			ctx, args.RoomName,
		)
	if err != nil {
		return fmt.Errorf("find room: %w", err)
	}

	if roomEntity != nil {
		return ErrUseCaseRoomCreateArgsRoomNameIsNotUnique
	}

	return nil
}

type UseCaseRoomCreateArgsValidatorGatewayUserFinder interface {
	UseCaseRoomCreateArgsValidatorGatewayFindUser(
		ctx context.Context,
		userID entity.UserID,
	) (
		*entity.User,
		error,
	)
}

type UseCaseRoomCreateArgsValidatorGatewayRoomFinder interface {
	UseCaseRoomCreateArgsValidatorGatewayFindRoom(
		ctx context.Context,
		roomName entity.RoomName,
	) (
		*entity.Room,
		error,
	)
}
