package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/exit/usecase/repository"

type Repository interface {
	repository.SessionDeactivator
}
