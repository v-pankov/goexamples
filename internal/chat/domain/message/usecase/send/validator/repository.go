package validator

import "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/validator/repository"

type Repository interface {
	repository.SessionFinder
}
