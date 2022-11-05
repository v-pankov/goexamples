package inmem

import (
	"context"
	"errors"
	"sync"

	"github.com/vdrpkv/goexamples/internal/chat/domain/message"
	"github.com/vdrpkv/goexamples/internal/chat/domain/session"

	messageSendUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/message/usecase/send/usecase"
	sessionLoginUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/login/usecase"
	sessionLogoutUsecase "github.com/vdrpkv/goexamples/internal/chat/domain/session/usecase/logout/usecase"
)

type MessageBus interface {
	SessionLoginUsecaseMessageBus() sessionLoginUsecase.MessageBus
	SessionLogoutUsecaseMessageBus() sessionLogoutUsecase.MessageBus
	MessageSendUsecaseMessageBus() messageSendUsecase.MessageBus
	Close() error
}

func New() MessageBus {
	return &messageBus{
		subs: make(map[session.ID]*messageBusSub),
	}
}

type messageBus struct {
	sync.Mutex
	subs    map[session.ID]*messageBusSub
	pending sync.WaitGroup
	closed  bool
}

type messageBusSub struct {
	writers  sync.WaitGroup
	messages chan *message.Entity
}

func (m *messageBus) SessionLoginUsecaseMessageBus() sessionLoginUsecase.MessageBus {
	return sessionLoginUsecaseMessageBusAdapter{m}
}

func (m *messageBus) SessionLogoutUsecaseMessageBus() sessionLogoutUsecase.MessageBus {
	return sessionLogoutUsecaseMessageBusAdapter{m}
}

func (m *messageBus) MessageSendUsecaseMessageBus() messageSendUsecase.MessageBus {
	return messageSendUsecaseMessageBusAdapter{m}
}

func (m *messageBus) Close() error {
	m.Lock()
	m.closed = true
	m.Unlock()

	m.pending.Wait()
	return nil
}

type sessionLoginUsecaseMessageBusAdapter struct {
	*messageBus
}

func (a sessionLoginUsecaseMessageBusAdapter) SubscribeSessionForNewMessages(
	ctx context.Context,
	sessionID session.ID,
) (
	<-chan *message.Entity,
	error,
) {
	var (
		val <-chan *message.Entity
		err = a.do(func() error {
			messages, err := a.subscribeSessionForNewMessages(ctx, sessionID)
			if err == nil {
				val = messages
			}
			return err
		})
	)
	return val, err
}

func (a sessionLoginUsecaseMessageBusAdapter) subscribeSessionForNewMessages(
	ctx context.Context,
	sessionID session.ID,
) (
	<-chan *message.Entity,
	error,
) {
	if sub := a.subs[sessionID]; sub != nil {
		return nil, errors.New("session is already subscribed")
	}

	sub := &messageBusSub{
		messages: make(chan *message.Entity),
	}

	a.subs[sessionID] = sub

	return sub.messages, nil
}

type sessionLogoutUsecaseMessageBusAdapter struct {
	*messageBus
}

func (a sessionLogoutUsecaseMessageBusAdapter) UnsubscribeSessionFromNewMessages(ctx context.Context, sessionID session.ID) error {
	return a.do(func() error {
		return a.unsubscribeSessionFromNewMessages(ctx, sessionID)
	})
}

func (a sessionLogoutUsecaseMessageBusAdapter) unsubscribeSessionFromNewMessages(_ context.Context, sessionID session.ID) error {
	sub := a.subs[sessionID]

	if sub == nil {
		return errors.New("session is not subscribed")
	}

	delete(a.subs, sessionID)

	a.pending.Add(1)
	go func() {
		sub.writers.Wait()
		a.pending.Done()
	}()

	return nil
}

type messageSendUsecaseMessageBusAdapter struct {
	*messageBus
}

func (a messageSendUsecaseMessageBusAdapter) BroadcastMessageToAllSessions(ctx context.Context, message *message.Entity) error {
	return a.do(func() error {
		return a.broadcastMessageToAllSessions(ctx, message)
	})
}

func (a messageSendUsecaseMessageBusAdapter) broadcastMessageToAllSessions(ctx context.Context, message *message.Entity) error {
	for _, sub := range a.subs {
		sub := sub
		sub.writers.Add(1)
		go func() {
			select {
			case <-ctx.Done():
				// nothing
			case sub.messages <- message:
				// nothing
			}
			sub.writers.Done()
		}()
	}
	return nil
}

func (m *messageBus) do(fn func() error) error {
	m.Lock()
	defer m.Unlock()

	if m.closed {
		return errors.New("message bus is closed")
	}

	return fn()
}
