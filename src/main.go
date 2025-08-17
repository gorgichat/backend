package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log/slog"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	_ =	r.Run()
}
