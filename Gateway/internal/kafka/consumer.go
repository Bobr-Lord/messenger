package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	TopicMessageSend      = "message.send"
	TopicMessageDelivered = "message.delivered"
)

type ConsumerKafka struct {
	reader *kafka.Reader
}

func NewConsumerMessage(brokers []string) *ConsumerKafka {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       TopicMessageDelivered,
		GroupID:     "message-service-group",
		StartOffset: kafka.FirstOffset,
		MinBytes:    1e3,  // 1KB
		MaxBytes:    10e6, // 10MB
		MaxWait:     1 * time.Second,
	})
	return &ConsumerKafka{
		reader: r,
	}
}

func (c *ConsumerKafka) Start(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			logrus.Errorf("Error reading message: %v", err)
			continue
		}
		logrus.Infof("Message received from topic %s: %s", m.Topic, string(m.Value))
	}
}

func (c *ConsumerKafka) Close() error {
	return c.reader.Close()
}
