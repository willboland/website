package message

import "time"

type Message struct {
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}

func New(author string, content string) Message {
	return Message{
		Author:    author,
		CreatedAt: time.Now(),
		Content:   content,
	}
}
