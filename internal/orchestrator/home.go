package orchestrator

import (
	"encoding/json"
	"github.com/willboland/website/internal/cache"
	"github.com/willboland/website/internal/message"
	"net/http"
)

func HomeGet(messages []message.Message) Response {
	if len(messages) == 0 {
		return Response{
			code: http.StatusOK,
			body: []byte(`No messages yet :/`),
		}
	}

	messagesJSON, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		return Response{
			code: http.StatusInternalServerError,
			body: []byte(`Something went wrong :(`),
		}
	}

	return Response{
		code: http.StatusOK,
		body: messagesJSON,
	}
}

func HomePost(r *http.Request, cache *cache.DestructiveCache[[]message.Message]) Response {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var m message.Message
	err := decoder.Decode(&m)
	if err != nil {
		return Response{
			code: http.StatusBadRequest,
			body: []byte(`{"error":"invalid request"}`),
		}
	}

	formattedMessage := message.New(m.Author, m.Content)
	messages, _ := cache.Get(m.Author)
	messages = append(messages, formattedMessage)
	cache.Set(formattedMessage.Author, messages)
	return Response{
		code: http.StatusCreated,
	}
}
