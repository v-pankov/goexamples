package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/exit/usecase/repository"

type Repository interface {
	repository.SessionDeactivator
}
