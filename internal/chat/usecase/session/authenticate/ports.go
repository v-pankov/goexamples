package authenticate

import "github.com/vdrpkv/goexamples/internal/chat/entity/session"

type Request struct {
	SessionID session.ID
}

type Response struct {
	SessionEntity *session.Entity
}
