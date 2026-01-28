package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "A simple web server using Gin framework",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Value:   8080,
				Usage:   "Port to run the web server on",
			},
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a"},
				Value:   "0.0.0.0",
				Usage:   "Address to bind the server to",
			},
			&cli.StringFlag{
				Name:    "env",
				Aliases: []string{"e"},
				Value:   "dev",
				Usage:   "Environment (dev|prod)",
			},
		},
		Action: func(c *cli.Context) error {
			port := c.Int("port")
			addr := c.String("addr")
			env := c.String("env")
			startServer(addr, port, env)
			return nil
		},
	}

	// Running the CLI app
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func startServer(addr string, port int, env string) {
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows) by default
	listenAddr := fmt.Sprintf("%s:%d", addr, port)
	log.Printf("Starting server on %s...", listenAddr)
	if listenAddr == ":" {
		log.Printf("Starting server on %s...", listenAddr)
	}

	err := r.Run(listenAddr)
	if err != nil {
		log.Fatalf("failed running Gin app: %s", err.Error())
	}
}
