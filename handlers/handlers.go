// TODO
package handlers

import (
	"context"
	"log"

	"github.com/Drynok/fhome_ha/handlers/feed"
	"github.com/Drynok/fhome_ha/handlers/history"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// NewRouterGroup creates new router group.
func NewRouterGroup(ctx context.Context, con *dig.Container, rtr *gin.Engine) (*gin.RouterGroup, error) {
	grp := rtr.Group("/v1")

	for _, err := range []error{
		con.Invoke(func(prm feed.Params) {
			// - accepts input data as JSON via an HTTP endpoint (/feed)
			// - shards the above data in such a way that the number of shards
			// does not exceed 5
			// - data from every shard is processed by exactly one worker from
			// a worker pool; the worker pool has the following
			// characteristics:
			// - any worker has a unique identifier
			// - no worker lives for more than 2 minutes (this means there
			// is a chance no worker is alive at some point)
			// - the initial number of workers is 3 and this number cannot
			// exceed 4
			// - processing the data means writing it to disk in batches
			// of 5 input items (in a file named using the worker
			// identifier)
			// - exposes an HTTP endpoint (/history) which returns a list of
			// worker identifiers and the number of processed messages for
			// each worker
			grp.GET("/feed", feed.NewHandler(ctx, &prm))
		}),

		con.Invoke(func(prm history.Params) {
			// - exposes an HTTP endpoint (/history) which returns a list of
			// worker identifiers and the number of processed messages for
			// each worker
			grp.GET("/history", gin.BasicAuth(gin.Accounts{
				"admin": "secret"}), history.NewHandler(ctx, &prm))
		}),

		// r.POST("/login", auth.LoginHandler)
	} {
		if err != nil {
			log.Println(err, "there is a problem creating a router group")
			return nil, err
		}
	}

	return grp, nil
}
