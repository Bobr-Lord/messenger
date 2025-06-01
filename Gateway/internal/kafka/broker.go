package kafka

import (
	"context"
	"gitlab.com/bobr-lord-messenger/gateway/internal/handler"
)

type ConsumerInterface interface {
	Start(context.Context)
	Close() error
}
type Consumer struct {
	Consumer ConsumerInterface
}

type ProducerInterface interface {
	Send(ctx context.Context, key, value []byte) error
	Close() error
}
type Producer struct {
	Producer ProducerInterface
}

func NewConsumer(addr []string, h *handler.Handler) *Consumer {
	return &Consumer{
		Consumer: NewConsumerMessage(addr, h),
	}
}
func NewProducer(addr []string) *Producer {
	return &Producer{
		Producer: NewProducerMessage(addr),
	}
}
