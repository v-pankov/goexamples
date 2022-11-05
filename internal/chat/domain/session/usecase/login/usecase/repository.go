package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login/usecase/repository"

type Repository interface {
	repository.UserCreatorFinder
	repository.ActiveSessionCreator
}
