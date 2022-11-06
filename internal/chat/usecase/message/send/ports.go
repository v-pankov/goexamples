package send

import "github.com/vdrpkv/goexamples/internal/chat/entity/session"

type Request struct {
	AuthorUserSessionID session.ID
	MessageText         string
}

type Response struct {
}
