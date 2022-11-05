package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter/usecase/repository"

type Repository interface {
	UserRepository
	SessionRepository
}

type UserRepository interface {
	repository.UserCreatorFinder
}

type SessionRepository interface {
	repository.ActiveSessionCreator
}
