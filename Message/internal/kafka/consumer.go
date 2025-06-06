package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gitlab.com/bobr-lord-messenger/message/internal/models"
	"gitlab.com/bobr-lord-messenger/message/internal/repository"
	"io"
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
			if err == io.EOF || errors.Is(err, io.EOF) {
				return
			}
			fmt.Println(err)
			logrus.Errorf("Error reading message: %v", err)
			continue
		}

		logrus.Infof("Message received from topic %s: %s", m.Topic, string(m.Value))
		var reqKafka *models.KafkaRequestMessage
		if err := json.Unmarshal(m.Value, &reqKafka); err != nil {
			logrus.Errorf("Error unmarshalling message: %v", err)
			continue
		}
		msg := &models.Message{
			ChatID:   reqKafka.ChatID,
			SenderID: reqKafka.UserID,
			Content:  reqKafka.Content,
		}
		id, err := c.repo.Message.Save(msg)
		if err != nil {
			logrus.Errorf("Error saving message: %v", err)
			continue
		}
		logrus.Infof("message saved with ID: %s", id)

		users, err := c.repo.Message.UsersSendMess(msg.ChatID, msg.SenderID)
		if err != nil {
			logrus.Errorf("Error sending message: %v", err)
			continue
		}
		for _, user := range *users {
			var req models.MessageDelivery
			fmt.Println(user)
			req.UserID = user
			req.ChatID = msg.ChatID
			req.Content = msg.Content
			reqJS, err := json.Marshal(req)
			if err != nil {
				logrus.Errorf("Error marshalling message: %v", err)
				continue
			}
			if err := c.prod.Producer.Send(context.Background(), []byte(id), reqJS); err != nil {
				logrus.Errorf("Error sending message: %v", err)
				continue
			}
			logrus.Infof("message sent in topic %v: %+v", TopicMessageDelivered, string(reqJS))
		}
	}
}

func (c *ConsumerMessage) Close() error {
	return c.reader.Close()
}
