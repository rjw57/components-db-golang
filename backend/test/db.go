package test

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

var TestingDbNotConfigured = errors.New("TESTING_DATABASE_DSN environment variable is not set")

// OpenTestingDatabase opens a new session onto a testing database. If the TESTING_DATABASE_DSN
// environment variable is set, that is used as the connection DSN otherwise an ephemeral database
// is created via docker.
//
// A new pointer to a new sql.DB instance is returned along with a function which must be called to
// close the instance.
func OpenTestingDatabase() (*sql.DB, func() error, error) {
	dsn, ok := os.LookupEnv("TESTING_DATABASE_DSN")
	if !ok {
		return nil, nil, TestingDbNotConfigured
	}

	log.Print("Opening testing database from dsn specified in TESTING_DATABASE_DSN")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}
	postgres.SetQueryLogger(logQuery)
	return db, (func() error { return nil }), nil
}

func logQuery(_ context.Context, queryInfo postgres.QueryInfo) {
	sql, args := queryInfo.Statement.Sql()
	log.Printf("- SQL: %s Args: %v \n", sql, args)
	// log.Printf("- Debug SQL: %s \n", queryInfo.Statement.DebugSql())
	log.Printf("- Rows processed: %d\n", queryInfo.RowsProcessed)
	log.Printf("- Duration %s\n", queryInfo.Duration.String())
	// log.Printf("- Execution error: %v\n", queryInfo.Err)
	// callerFile, callerLine, callerFunction := queryInfo.Caller()
	// log.Printf("- Caller file: %s, line: %d, function: %s\n", callerFile, callerLine, callerFunction)
}
