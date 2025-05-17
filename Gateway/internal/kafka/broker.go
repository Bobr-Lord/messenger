package kafka

import "context"

type Producer interface {
	Send(ctx context.Context, key, value []byte) error
	Close() error
}

type Consumer interface {
	Start(ctx context.Context)
	Close() error
}

type BrokerProducer struct {
	Producer Producer
}
type BrokerConsumer struct {
	Consumer Consumer
}

func NewBrokerProducer(addr []string) *BrokerProducer {
	return &BrokerProducer{
		Producer: NewProducerMessage(addr),
	}
}
func NewBrokerConsumer(addr []string) *BrokerConsumer {
	return &BrokerConsumer{
		Consumer: NewConsumerMessage(addr),
	}
}
