package main

import (
	"os"
	"flag"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log/slog"
)

var (
	host = flag.String("host", "127.0.0.1", "Gorgi Chat host")
	port = flag.Int("port", 8080, "Gorgi Chat running port")

	dbHost     = flag.String("db-host", "localhost", "Database host")
	dbPort     = flag.Int("db-port", 3306, "Database port")
	dbUser     = flag.String("db-user", "root", "Database user")
	dbPassword = flag.String("db-password", "", "Database password")
	dbName     = flag.String("db-name", "gorgi_chat", "Database name")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	if val := os.Getenv("HOST"); val != "" {
		*host = val
	}
	if val := os.Getenv("PORT"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			*port = parsed
		}
	}
	if val := os.Getenv("DB_HOST"); val != "" {
		*dbHost = val
	}
	if val := os.Getenv("DB_PORT"); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil {
			*dbPort = parsed
		}
	}
	if val := os.Getenv("DB_USER"); val != "" {
		*dbUser = val
	}
	if val := os.Getenv("DB_PASSWORD"); val != "" {
		*dbPassword = val
	}
	if val := os.Getenv("DB_NAME"); val != "" {
		*dbName = val
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	_ =	r.Run()
}
