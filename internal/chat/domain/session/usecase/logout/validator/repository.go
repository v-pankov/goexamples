package validator

import "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/validator/repository"

type Repository interface {
	repository.SessionFinder
}
