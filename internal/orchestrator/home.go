package orchestrator

import (
	"bytes"
	"encoding/json"
	"github.com/willboland/simcache"
	"github.com/willboland/website/internal/message"
	"net/http"
	"strings"
	"time"
)

func HomeGet(messages []message.Message) (int, []byte) {
	if len(messages) == 0 {
		return http.StatusNoContent, []byte(`No messages yet :/`)
	}

	var builder strings.Builder
	for _, m := range messages {
		builder.WriteString(m.Author + " on " + m.CreatedAt.Format(time.DateTime) + " wrote: ")
		builder.WriteString(m.Content + "\n\n")
	}

	body := bytes.NewBufferString(builder.String()).Bytes()
	return http.StatusOK, body
}

func HomePost(r *http.Request, cache *simcache.Cache[message.Message]) (int, []byte) {
	type request struct {
		Author  string
		Message string
	}

	type response struct {
		Error string `json:"error,omitempty"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var requestBody request
	err := decoder.Decode(&requestBody)
	if err != nil {
		b, _ := json.Marshal(response{Error: err.Error()})
		return http.StatusBadRequest, b
	}

	formattedMessage := message.New(requestBody.Author, requestBody.Message)
	key := formattedMessage.Author + time.Now().Format(time.StampNano)
	added := cache.Add(key, formattedMessage)
	if !added {
		b, _ := json.Marshal(response{Error: "duplicate request occurred at same time, try again"})
		return http.StatusConflict, b
	}

	return http.StatusCreated, nil
}
