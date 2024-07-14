package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const ctxTransactionKey = "gormTx"

func Tx(ctx *gin.Context) *gorm.DB {
	tx, ok := ctx.Get(ctxTransactionKey)
	if !ok {
		return nil
	}
	return tx.(*gorm.DB)
}

func Database(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		db.WithContext(c).Transaction(func(tx *gorm.DB) error {
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

			return nil
		})
	}
}
