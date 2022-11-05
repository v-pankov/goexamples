package send

import "github.com/vdrpkv/goexamples/internal/chat/domain/session"

type Args struct {
	AuthorUserSessionID session.ID
	MessageText         string
}

type Result struct {
}
