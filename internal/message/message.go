package message

import (
	"strings"
	"time"
)

type Message struct {
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// New returns a Message with cleansed fields and CreatedAt as UTC.
func New(author string, content string) Message {
	return Message{
		Author:    strings.TrimSpace(author),
		Content:   strings.TrimSpace(content),
		CreatedAt: time.Now().UTC(),
	}
}
