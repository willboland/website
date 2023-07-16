package server

import (
	"github.com/willboland/website/internal/orchestrator"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	var code int
	var body []byte

	switch r.Method {
	case http.MethodGet:
		code, body = orchestrator.HomeGet(s.messageCache.Values())
	case http.MethodPost:
		code, body = orchestrator.HomePost(r, s.messageCache)
	default:
		http.Error(w, `{"error": "method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(body)
}
