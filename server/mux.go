package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ServeMux struct {
	mux *mux.Router
}

func NewServeMux() (*ServeMux, error) {
	serveMux := ServeMux{}

	serveMux.mux = mux.NewRouter()

	return &serveMux, nil
}

func (s *ServeMux) AddRoute(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(path, handler)
}
