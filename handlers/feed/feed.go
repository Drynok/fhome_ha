// Package status creates HTTP handler for example endpoint. And serves and status health check.
package feed

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Drynok/fhome_ha/packages/sharder"
	"github.com/gin-gonic/gin"
)

type Params struct {
	shr sharder.JsonSharder
}

type Model struct {
	Value     time.Time `json:"value" binding:"min=1,max=255` // time format is RFC 3339 (2019-10-12T07:20:50.52Z) for incoming requests
	Timestamp []string  `json:"timestamp" binding:"min=1,max=255"`
}

// NewHandler creates new HTTP handler.
func NewHandler(ctx context.Context, p *Params) gin.HandlerFunc {
	return func(c *gin.Context) {
		m := new(Model)

		if err := c.ShouldBind(m); err != nil && err != io.EOF {
			log.Fatalln(err, "Unprocessable entity")
			return
		}

		if err := c.ShouldBindJSON(m); err != nil && err != io.EOF {
			log.Fatalln(err, "Unprocessable JSON")
			return
		}

		c.Status(http.StatusOK)
	}
}
