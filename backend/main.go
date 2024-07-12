package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/api"
	"github.com/rjw57/components-db-golang/backend/middleware/logger"
)

func main() {
	// If we're not in release mode, enable the log pretty-printer.
	if gin.Mode() != gin.ReleaseMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Create a gin Engine and register API routes with it.
	r := NewGinEngine()
	api.RegisterHandlers(r, api.NewServer())

	// Serve HTTP until the world ends.
	log.Fatal().Err(NewHttpServer(r).ListenAndServe())
}

// NewGinEngine constructs a new gin.Engine instance with our desired middleware added.
func NewGinEngine() *gin.Engine {
	// Configure gin and add any required middleware.
	r := gin.New()
	r.Use(logger.DefaultStructuredLogger())
	r.Use(gin.Recovery())
	return r
}

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
