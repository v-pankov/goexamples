package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type UseCaseMessageReceive interface {
	DoUseCaseMessageReceive(
		ctx context.Context,
		args *UseCaseMessageReceiveArgs,
	) (
		*UseCaseMessageReceiveResult,
		error,
	)
}

type UseCaseMessageReceiveArgs struct {
	SessionID entity.SessionID
	Message   *entity.Message
}

type UseCaseMessageReceiveResult struct {
}

func NewUseCaseMessageReceive(
	gatewaySessionMessageDeliverer UseCaseMessageReceiveGatewaySessionMessageDeliverer,
) UseCaseMessageReceive {
	return useCaseMessageReceive{
		gatewaySessionMessageDeliverer: gatewaySessionMessageDeliverer,
	}
}

type UseCaseMessageReceiveGatewaySessionMessageDeliverer interface {
	UseCaseMessageReceiveGatewayDeliverMessageToSession(
		ctx context.Context,
		sessionID entity.SessionID,
		message *entity.Message,
	) error
}

type useCaseMessageReceive struct {
	gatewaySessionMessageDeliverer UseCaseMessageReceiveGatewaySessionMessageDeliverer
}

func (uc useCaseMessageReceive) DoUseCaseMessageReceive(
	ctx context.Context,
	args *UseCaseMessageReceiveArgs,
) (
	*UseCaseMessageReceiveResult,
	error,
) {
	err := uc.
		gatewaySessionMessageDeliverer.
		UseCaseMessageReceiveGatewayDeliverMessageToSession(
			ctx,
			args.SessionID,
			args.Message,
		)
	if err != nil {
		return nil, fmt.Errorf("deliver message to session: %w", err)
	}

	return &UseCaseMessageReceiveResult{}, nil
}

type UseCaseMessageReceiveArgsValidator interface {
	ValidateUseCaseMessageReceiveArgs(ctx context.Context, args *UseCaseMessageReceiveArgs) error
}

var (
	ErrUseCaseMessageReceiveArgsSessionDoesNotExist = errors.New("session does not exist")
)

func NewUseCaseMessageReceiveArgsValidator(
	gatewayUserFinder UseCaseMessageReceiveArgsValidatorGatewaySessionFinder,
) UseCaseMessageReceiveArgsValidator {
	return useCaseMessageReceiveArgsValidator{
		gatewaySessionFinder: gatewayUserFinder,
	}
}

type useCaseMessageReceiveArgsValidator struct {
	gatewaySessionFinder UseCaseMessageReceiveArgsValidatorGatewaySessionFinder
}

func (v useCaseMessageReceiveArgsValidator) ValidateUseCaseMessageReceiveArgs(
	ctx context.Context, args *UseCaseMessageReceiveArgs,
) error {
	if err := args.SessionID.Validate(); err != nil {
		return fmt.Errorf("session id: %w", err)
	}

	if err := args.Message.Validate(); err != nil {
		return fmt.Errorf("message: %w", err)
	}

	sessionEntity, err := v.
		gatewaySessionFinder.
		UseCaseMessageReceiveArgsValidatorGatewayFindSession(
			ctx, args.SessionID,
		)
	if err != nil {
		return fmt.Errorf("find session: %w", err)
	}

	if sessionEntity == nil {
		return ErrUseCaseMessageReceiveArgsSessionDoesNotExist
	}

	return nil
}

type UseCaseMessageReceiveArgsValidatorGatewaySessionFinder interface {
	UseCaseMessageReceiveArgsValidatorGatewayFindSession(
		ctx context.Context, sessionID entity.SessionID,
	) (
		*entity.Session,
		error,
	)
}
