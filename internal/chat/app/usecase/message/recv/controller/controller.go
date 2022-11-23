package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vdrpkv/goexamples/internal/chat/app/controller"
)

type Controller struct {
	Viewer ModelViewer
}

var _ controller.Controller = Controller{}

func (c Controller) HandleMessage(ctx context.Context, message []byte) error {
	var input struct {
		MessageID       int64  `json:"id"`
		MessageContents []byte `json:"contents"`
		CreatedAt       int64  `json:"created_at"`
	}

	if err := json.Unmarshal(message, &input); err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	if err := c.Viewer.ViewModel(ctx, &ViewModel{
		MessageID:       input.MessageID,
		MessageContents: input.MessageContents,
		CreatedAt:       input.CreatedAt,
	}); err != nil {
		return fmt.Errorf("view: %w", err)
	}

	return nil
}

type ViewModel struct {
	MessageID       int64
	MessageContents []byte
	CreatedAt       int64
}

type ModelViewer interface {
	ViewModel(context.Context, *ViewModel) error
}
