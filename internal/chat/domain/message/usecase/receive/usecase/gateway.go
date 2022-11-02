package usecase

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"
)

type GatewaySessionMessageDeliverer interface {
	GatewayDeliverMessageToSession(
		ctx context.Context,
		sessionID session.ID,
		message *message.Entity,
	) error
}
