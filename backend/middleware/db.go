package middleware

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

const ctxTransactionKey = "databaseTransaction"

func Tx(ctx *gin.Context) *sql.Tx {
	tx, ok := ctx.Get(ctxTransactionKey)
	if !ok {
		return nil
	}
	return tx.(*sql.Tx)
}

func Database(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := db.Begin()
		if err != nil {
			log.Error().Err(err).Msg("Failed to start database transaction")
			c.AbortWithStatus(500)
		}

		c.Set(ctxTransactionKey, tx)
		c.Next()
		c.Set(ctxTransactionKey, nil)

		if err := recover(); err != nil {
			tx.Rollback()
		} else if c.Writer.Status() == http.StatusInternalServerError {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
