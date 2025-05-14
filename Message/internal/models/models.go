package models

type Message struct {
	ChatID   string `json:"chat_id" db:"chat_id"`
	SenderID string `json:"sender_id" db:"sender_id"`
	Content  string `json:"content" db:"content"`
	SentAt   int64  `json:"sent_at" db:"sent_at"`
}
