package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

// NewHttpServer constructs a new http.Server instance respecting the PORT and HOST environment
// variables if set.
func NewHttpServer(handler http.Handler) *http.Server {
	bind_host := "0.0.0.0"
	bind_port := "8000"

	if v, ok := os.LookupEnv("HOST"); ok {
		bind_host = v
	}

	if v, ok := os.LookupEnv("PORT"); ok {
		bind_port = v
	}

	bind_addr := fmt.Sprintf("%v:%v", bind_host, bind_port)
	log.Info().Str("bind_addr", bind_addr).Msg("Starting server")

	return &http.Server{
		Handler: handler,
		Addr:    bind_addr,
	}
}
