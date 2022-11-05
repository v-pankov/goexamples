package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/enter/usecase/repository"

type Repository interface {
	repository.UserCreatorFinder
	repository.ActiveSessionCreator
}
