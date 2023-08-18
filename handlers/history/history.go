package history

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Parameters dependency injection for the handler.
type Params struct {
	dig.In
}

// NewHandler creates new example HTTP handler.
func NewHandler(ctx context.Context, p *Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		for {
			c.SSEvent("Event", "example event message...")
			c.Writer.Flush()
			time.Sleep(1 * time.Second)

			// Break when request context is cancelled.
			if err := c.Request.Context().Err(); err == context.Canceled {
				log.Println("API request's context cancelled.")
				break
			}

			// Break when the server's context is cancelled.
			if err := ctx.Err(); err == context.Canceled {
				log.Println("API server's context cancelled.")
				break
			}
		}
	}
}
