package ui

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = 25400

func checkPort() bool {

	return false
}

func web() *http.Server {
	if checkPort() {
		return nil
	}
	r := gin.New()
	r.Static("/", "./public")
	r.POST("/hello", func(c *gin.Context) {
		c.Status(200)
	})
	srv := &http.Server{
		Addr:    ":25400",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return srv
}
