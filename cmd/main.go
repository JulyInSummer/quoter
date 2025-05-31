package main

import (
	"flag"
	"github.com/JulyInSummer/quoter_app/internal/service"
	"github.com/JulyInSummer/quoter_app/internal/storage/postgres"
	"github.com/JulyInSummer/quoter_app/internal/transport/http"
	"log/slog"
	"os"
)

func main() {
	level := getLogLevel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	logger.Info("Starting server")

	// connect to a database
	db := postgres.NewConn()

	// create postgres storage
	strg := postgres.NewQuoteStorage(db, logger)

	// create the service layer
	srvc := service.NewQuoteService(logger, strg)

	// create handler
	handler := http.NewHandler(logger, srvc)

	// create server
	server := http.NewServer(":8080", handler)

	server.Run()
}

func getLogLevel() slog.Level {
	var (
		levelInput string
		level      slog.Level
	)
	flag.StringVar(&levelInput, "log-level", "info", "Setting the logging level of the application.")
	flag.Parse()

	switch levelInput {
	case "debug":
		level = -4
	default:
		level = 0
	}

	return level
}
