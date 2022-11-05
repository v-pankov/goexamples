package validator

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/exit/validator/repository"

type Repository interface {
	SessionRepository
}

type SessionRepository interface {
	repository.SessionFinder
}
