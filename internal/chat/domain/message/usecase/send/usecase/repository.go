package usecase

import "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/usecase/repository"

type Repository interface {
	MessageRepository
}

type MessageRepository interface {
	repository.MessageCreator
}
