package main

import (
	"github.com/joho/godotenv"
	"github.com/lucasszmt/transfeera-challenge/app"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/lucasszmt/transfeera-challenge/infra/db"
	"github.com/lucasszmt/transfeera-challenge/infra/log"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	logger := log.NewLogger(zerolog.DebugLevel)

	// Init DB connection
	dbConn := db.Must(db.NewPostgresConn(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))

	// Init repositories
	receiverRepo := db.NewReceiver(dbConn)

	// Init services
	receiverService := receiver.NewService(&logger, receiverRepo)

	server := app.NewServer(receiverService)
	server.Run()
}
