package main

import (
	"os"
	"flag"
	"time"
	"strconv"
	"net/http"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/gorgichat/backend/pkg/routes"
	"github.com/gorgichat/backend/pkg/database"
	// "github.com/gorgichat/backend/pkg/localization"
)

var (
	host       = flag.String("host", "127.0.0.1", "Gorgi Chat host")
	port       = flag.Int("port", 8080, "Gorgi Chat running port")
	// language   = flag.String("language", "en", "Language for the chat")

	dbHost     = flag.String("db-host", "127.0.0.1", "Database host")
	dbPort     = flag.Int("db-port", 3306, "Database port")
	dbUser     = flag.String("db-user", "gorgi", "Database user")
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

	// Connect to the database
	err = database.ConnectDatabase(*dbHost, *dbPort, *dbUser, *dbPassword, *dbName)
	if err != nil {
		slog.Error("Error connecting to database: ", "err", err)
		return
	}
	// Migrate the database
	database.MigrateDatabase(database.GetDB())

	// Initialize i18n
	// bundle := i18n.NewBundle(language.English)

	// Initialize Gin
	r := gin.Default()
	routes.SetupRoutes(r)

	s := &http.Server{
		Addr:         *host + ":" + strconv.Itoa(*port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start the server
	slog.Info("Starting server at", "url", "http://"+ *host + ":" + strconv.Itoa(*port))
	if err := s.ListenAndServe(); err != nil {
		slog.Error("Error starting server: ", "err", err)
		return
	}
}
