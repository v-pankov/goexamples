package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/exit/usecase/repository"

type Repository interface {
	SessionRepository
}

type SessionRepository interface {
	repository.SessionDeactivator
}
