package server

import (
	"github.com/willboland/website/internal/cache"
	"github.com/willboland/website/internal/message"
	"net/http"
)

type Server struct {
	Router       *http.ServeMux
	messageCache *cache.DestructiveCache[[]message.Message]
}

// NewServer returns a new Server with routes pre-configured.
func NewServer() *Server {
	s := &Server{Router: http.NewServeMux()}
	s.SetupRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}
