package history

import (
	"context"
	"net/http"

	wp "github.com/Drynok/fhome_ha/packages/workerpool"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Parameters dependency injection for the handler.
type Params struct {
	dig.In
}

// NewHandler creates new example HTTP handler.
// exposes an HTTP endpoint (/history) which returns a list of
// worker identifiers and the number of processed messages for
// each worker
func NewHandler(ctx context.Context, p *Params, wrp wp.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, wrp.GetStats())
	}
}
