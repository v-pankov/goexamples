package gateways

import (
	"context"

	"github.com/vdrpkv/goexamples/internal/chat/entity"
)

type Repository interface {
	CreateMessage(context.Context, entity.MessageContents) (*entity.Message, error)
}
