package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/message/internal/models"
	"gitlab.com/bobr-lord-messenger/message/internal/repository"
	"time"
)

const (
	TopicMessageSend      = "message.send"
	TopicMessageDelivered = "message.delivered"
)

type ConsumerMessage struct {
	reader *kafka.Reader
	repo   *repository.Repository
	prod   *ProducerKafka
}

func NewConsumerMessage(brokers []string, repo *repository.Repository, prod *ProducerKafka) *ConsumerMessage {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       TopicMessageSend,
		GroupID:     "message-service-group",
		StartOffset: kafka.FirstOffset,
		MinBytes:    1e3,  // 1KB
		MaxBytes:    10e6, // 10MB
		MaxWait:     1 * time.Second,
	})
	return &ConsumerMessage{
		reader: r,
		repo:   repo,
		prod:   prod,
	}
}

func (c *ConsumerMessage) Start(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			logrus.Errorf("Error reading message: %v", err)
			continue
		}

		logrus.Infof("Message received from topic %s: %s", m.Topic, string(m.Value))
		var msg *models.Message
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			logrus.Errorf("Error unmarshalling message: %v", err)
			continue
		}
		logrus.Infof("message received: %+v", msg)
		id, err := c.repo.Message.Save(msg)
		if err != nil {
			logrus.Errorf("Error saving message: %v", err)
			continue
		}
		logrus.Infof("message saved with ID: %s", id)
		if err := c.prod.Producer.Send(context.Background(), []byte(msg.ChatID), m.Value); err != nil {
			logrus.Errorf("Error sending message: %v", err)
			continue
		}
		logrus.Info("message sent")
	}
}

func (c *ConsumerMessage) Close() error {
	return c.reader.Close()
}
