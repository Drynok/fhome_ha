package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Drynok/fhome_ha/env"
	"github.com/Drynok/fhome_ha/handlers"
	"github.com/Drynok/fhome_ha/packages/container"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctr, err := container.New()

	if err != nil {
		log.Panic(err)
	}

	ctx := context.Background()

	err = ctr.Invoke(func(env *env.Environment) {
		gin.SetMode(env.ServerMode)

		rtr := gin.New()
		rtr.Use(gin.Recovery())
		rtr.Use(gin.Logger())

		rtr.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "*"
			},
			MaxAge: 12 * time.Hour,
		}))

		rtr.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})

		if _, err := handlers.NewRouterGroup(ctx, ctr, rtr); err != nil {
			log.Panic(err)
		}

		srv := &http.Server{
			Addr:              fmt.Sprintf(":%s", env.ServerPort),
			Handler:           h2c.NewHandler(rtr, &http2.Server{}),
			ReadHeaderTimeout: time.Second * 60,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Panic(err)
			}
		}()
	})

	if err != nil {
		log.Panic(err)
	}
}
