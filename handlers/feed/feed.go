// Package status creates HTTP handler for example endpoint. And serves and status health check.
package feed

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Drynok/fhome_ha/packages/sharder"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type Params struct {
	dig.In
}

// Model for input json.
type Model struct {
	Value     string `json:"value" binding:"required,min=1,max=255`
	Timestamp string `json:"timestamp" binding:"required,min=1,max=255"`
}

// NewHandler creates new HTTP handler.
func NewHandler(ctx context.Context, p *Params, wpl wp.WorkerPool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data []Model

		jsd, err := io.ReadAll(c.Request.Body)

		if err != nil {
			log.Fatalln(err, "Unprocessable JSON")
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		// Validation.
		err = json.Unmarshal(jsd, &data)
		if err != nil {
			log.Fatalln(err, "Unprocessable JSON")
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		sds, err := sharder.Shard(data, 5)

		if err != nil {
			log.Fatalln(err, "Unable to shard incoming JSON")
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		for _, shd := range sds {

		}

		c.Status(http.StatusOK)
	}
}
