package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	cli "github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "gohair",
		Usage: "A simple web server using Gin framework",
		Action: func(c *cli.Context) error {
			startServer()
			return nil
		},
	}

	err := app.Run([]string{"gohair"})
	if err != nil {
		panic(err)
	}
}

func startServer() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
