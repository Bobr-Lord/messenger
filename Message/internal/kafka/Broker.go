package kafka

import (
	"context"
	"gitlab.com/bobr-lord-messenger/message/internal/repository"
)

type Consumer interface {
	Start(ctx context.Context)
	Close() error
}
type Producer interface {
}

type Broker struct {
	Consumer Consumer
}

func NewBreaker(repo *repository.Repository, addr []string) *Broker {
	return &Broker{
		Consumer: NewConsumerMessage(addr, repo),
	}
}
