package server

import (
	"github.com/willboland/simcache"
	"github.com/willboland/website/internal/message"
	"net/http"
	"time"
)

type Server struct {
	Router       *http.ServeMux
	messageCache *simcache.Cache[message.Message]
}

// NewServer returns a new Server with routes and dependencies pre-configured.
func NewServer() *Server {
	s := &Server{
		Router:       http.NewServeMux(),
		messageCache: simcache.New[message.Message](time.Hour * 24),
	}
	s.SetupRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}
