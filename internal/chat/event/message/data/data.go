package data

import "github.com/vdrpkv/goexamples/internal/chat/event"

type Created struct {
	MessageContents event.MessageContents
}

type Updated struct {
	MessageContents event.MessageContents
}

type Deleted struct {
}
