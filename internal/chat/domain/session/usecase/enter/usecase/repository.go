package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/enter/usecase/repository"

type Repository interface {
	repository.UserCreatorFinder
	repository.ActiveSessionCreator
}
