package models

type Message struct {
	ChatID   string `json:"chat_id" db:"chat_id"`
	SenderID string `json:"sender_id" db:"sender_id"`
	Content  string `json:"content" db:"content"`
	SentAt   int64  `json:"sent_at" db:"sent_at"`
}

type MessageDelivery struct {
	UserID  string `json:"user_id" db:"user_id"`
	ChatID  string `json:"chat_id" db:"chat_id"`
	Content string `json:"content" db:"content"`
}
type GetMessagesFromChatRequest struct {
	ChatID string `json:"chat_id" db:"chat_id"`
}
type GetMessagesFromChatResponse struct {
	Messages []*Message `json:"messages"`
}

type GetUserMessageRequest struct {
	UserID string `json:"user_id" db:"user_id"`
}
type GetUserMessageResponse struct {
	Messages []*Message `json:"message"`
}

type KafkaRequestMessage struct {
	UserID  string `json:"user_id" db:"user_id"`
	ChatID  string `json:"chat_id" db:"chat_id"`
	Content string `json:"content" db:"content"`
}
