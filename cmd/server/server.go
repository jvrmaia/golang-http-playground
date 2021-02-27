package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jvrmaia/golang-http-playground/logger"
)

func ok(w http.ResponseWriter, r *http.Request) {
	log := logger.New(false)
	log.Info().Msg("ok")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	log := logger.New(false)
	log.Info().Msg("timeout")
	time.Sleep(1 * time.Minute)
	w.WriteHeader(http.StatusGatewayTimeout)
	fmt.Fprint(w, "timeout from server")
}

func slow(w http.ResponseWriter, r *http.Request) {
	log := logger.New(false)
	log.Info().Msg("slow")
	time.Sleep(2 * time.Minute)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "slow")
}

func main() {
	log := logger.New(false)
	h := mux.NewRouter()
	h.HandleFunc("/slow", slow)
	h.HandleFunc("/timeout", timeout)
	h.HandleFunc("/", ok)
	s := &http.Server{
		Addr:         ":8080",
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("server startup failed")
	}
}
