package orchestrator

import "net/http"

type Response struct {
	code int
	body []byte
}

func (r Response) Write(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	_, _ = w.Write(r.body)
}
