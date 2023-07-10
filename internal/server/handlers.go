package server

import (
	"github.com/willboland/website/internal/cache"
	"github.com/willboland/website/internal/message"
	"github.com/willboland/website/internal/orchestrator"
	"net/http"
	"time"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	if s.messageCache == nil {
		s.messageCache = cache.NewDestructiveCache[[]message.Message](24 * time.Hour)
		go s.messageCache.EnqueueDestruction()
	}

	var messages []message.Message
	for _, authorsMessages := range s.messageCache.Values() {
		messages = append(messages, authorsMessages...)
	}

	switch r.Method {
	case http.MethodGet:
		orchestrator.HomeGet(messages).Write(w)
	case http.MethodPost:
		orchestrator.HomePost(r, s.messageCache).Write(w)
	default:
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
	}
}
