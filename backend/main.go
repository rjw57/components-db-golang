//go:generate docker compose run --no-TTY --rm generate-jet-schema
package main

import (
	"context"
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-jet/jet/v2/postgres"
	oapivalidate "github.com/oapi-codegen/gin-middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/rjw57/components-db-golang/backend/api"
	"github.com/rjw57/components-db-golang/backend/db/schema/components/public/table"
	"github.com/rjw57/components-db-golang/backend/middleware"
)

func main() {
	// If we're not in release mode, enable the log pretty-printer.
	if gin.Mode() != gin.ReleaseMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		postgres.SetQueryLogger(queryLogger)
	}

	// Create a database session.
	db, err := OpenDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening database")
	}

	// Create a gin Engine and register API routes with it.
	r := NewGinEngine(db)
	api.RegisterHandlers(r, api.NewServer())

	// Configure the schema name for the database.
	table.UseSchema("public")

	// Serve HTTP until the world ends.
	log.Fatal().Err(NewHttpServer(r).ListenAndServe())
}

// NewGinEngine constructs a new gin.Engine instance with our desired middleware added.
func NewGinEngine(db *sql.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(middleware.Database(db))

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting API specification")
	}
	r.Use(oapivalidate.OapiRequestValidator(swagger))

	return r
}

func queryLogger(ctx context.Context, queryInfo postgres.QueryInfo) {
	sql, args := queryInfo.Statement.Sql()
	log.Info().
		Str("sql", sql).
		Any("args", args).
		Int64("rows_processed", queryInfo.RowsProcessed).
		Any("duration_ns", queryInfo.Duration).
		Str("duration", queryInfo.Duration.String()).
		Msg("Executed SQL query")
}
