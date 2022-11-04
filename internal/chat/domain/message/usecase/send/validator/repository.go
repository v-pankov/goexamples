package validator

import "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/validator/repository"

type Repository interface {
	UserRepository
}

type UserRepository interface {
	repository.UserFinder
}
