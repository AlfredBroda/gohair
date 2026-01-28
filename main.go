package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cli "github.com/urfave/cli/v2"

	"github.com/AlfredBroda/gohair/model"
)

const (
	DB_PASSWORD = "123pass"
	DB_USER     = "root"
	DB_ADDR     = "127.0.0.1"
	DB_PORT     = 3306
	DB_NAME     = "rhair"
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
			&cli.StringFlag{
				Name:    "dbuser",
				EnvVars: []string{"DB_USER"},
				Value:   DB_USER,
				Usage:   "Database user",
			},
			&cli.StringFlag{
				Name:    "dbpass",
				EnvVars: []string{"DB_PASSWORD"},
				Value:   DB_PASSWORD,
				Usage:   "Database password",
			},
			&cli.StringFlag{
				Name:    "dbaddr",
				EnvVars: []string{"DB_ADDRESS"},
				Value:   DB_ADDR,
				Usage:   "Database address",
			},
			&cli.IntFlag{
				Name:    "dbport",
				EnvVars: []string{"DB_PORT"},
				Value:   DB_PORT,
				Usage:   "Database port",
			},
			&cli.StringFlag{
				Name:    "dbname",
				EnvVars: []string{"DB_NAME"},
				Value:   DB_NAME,
				Usage:   "Database name",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "Run database migrations",
				Action: func(c *cli.Context) error {
					dbConfig := model.DBConfig{
						DBUser: c.String("dbuser"),
						DBPass: c.String("dbpass"),
						DBAddr: c.String("dbaddr"),
						DBPort: c.Int("dbport"),
						DBName: c.String("dbname"),
					}
					dialector := model.ConfigureMySQL(dbConfig)
					err := model.Migrate(dialector)
					if err != nil {
						fmt.Println("Database migration failed:", err)
						return err
					}
					fmt.Println("Database migration completed.")
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			startServer(GinConfig{
				Env:  c.String("env"),
				Port: c.Int("port"),
				Addr: c.String("addr"),
			}, model.DBConfig{
				DBUser: c.String("dbuser"),
				DBPass: c.String("dbpass"),
				DBAddr: c.String("dbaddr"),
				DBPort: c.Int("dbport"),
				DBName: c.String("dbname"),
			})
			return nil
		},
	}

	// Running the CLI app
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

type GinConfig struct {
	Env  string
	Port int
	Addr string
}

func startServer(ginConfig GinConfig, dbConfig model.DBConfig) {
	if ginConfig.Env == "prod" {
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

	dialector := model.ConfigureMySQL(dbConfig)
	r.GET("/status/db", func(c *gin.Context) {
		// Return JSON response

		_, err := model.InitDB(dialector)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to connect to database: %s", err.Error()),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Database connected successfully",
		})
	})

	r.GET("/a/:slug", func(c *gin.Context) {
		slug := c.Param("slug")
		article, err := model.GetArticleBySlug(dialector, slug)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Article not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"article": article,
		})
	})

	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows) by default
	listenAddr := fmt.Sprintf("%s:%d", ginConfig.Addr, ginConfig.Port)
	log.Printf("Starting server on %s...", listenAddr)
	if listenAddr == ":" {
		log.Printf("Starting server on %s...", listenAddr)
	}

	err := r.Run(listenAddr)
	if err != nil {
		log.Fatalf("failed running Gin app: %s", err.Error())
	}
}
