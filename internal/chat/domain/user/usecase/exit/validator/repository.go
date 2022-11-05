package validator

import "github.com/vdrpkv/goexamples/internal/chat/domain/user/usecase/exit/validator/repository"

type Repository interface {
	repository.SessionFinder
}
