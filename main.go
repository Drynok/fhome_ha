package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Drynok/fhome_ha/env"
	"github.com/Drynok/fhome_ha/handlers"
	"github.com/Drynok/fhome_ha/packages/container"
	wp "github.com/Drynok/fhome_ha/packages/workerpool"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Set better loging.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctr, err := container.New()

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Dependency injection.
	err = ctr.Invoke(func(env *env.Environment) {
		// Init data channel.
		dch := make(chan string)
		defer close(dch)

		// Init pool of workers.
		wrp := wp.NewPool(3, 5)

		gin.SetMode(env.ServerMode)
		rtr := gin.New()

		//
		rtr.Use(gin.Recovery())
		//
		rtr.Use(gin.Logger())

		// CORS policies.
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

		// 404 route.
		rtr.NoRoute(func(c *gin.Context) {
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		})

		// Init group of endpoints.
		if _, err := handlers.NewRouterGroup(ctx, ctr, rtr, *wrp); err != nil {
			log.Panic(err)
		}

		// Init server.
		srv := &http.Server{
			Addr:              fmt.Sprintf(":%s", env.ServerPort),
			Handler:           h2c.NewHandler(rtr, &http2.Server{}),
			ReadHeaderTimeout: time.Second * 60,
		}

		// Run server.
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Panic(err)
			}
		}()

		// Gracefully shutdown of server.
		quit := make(chan os.Signal)
		// kill signals.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		if err := srv.Shutdown(ctx); err != nil {
			log.Panic("Server Shutdown:", err)
		}

		// catching ctx.Done().
		select {
		case <-ctx.Done():
			log.Println("timeout of 5 seconds.")
		}
		log.Println("Server exiting")
	})

	if err != nil {
		log.Panic(err)
	}
}
