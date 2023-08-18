// Package status creates HTTP handler for example endpoint. And serves and status health check.
package feed

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Timestamp string `json:`
	Value     string `json`
}

// NewHandler creates new HTTP handler.
func NewHandler(ctx context.Context, p *Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
