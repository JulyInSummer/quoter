package main

import (
	"github.com/JulyInSummer/quoter_app/internal/service"
	"github.com/JulyInSummer/quoter_app/internal/storage/postgres"
	"github.com/JulyInSummer/quoter_app/internal/transport/http"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting server")

	// connect to a database
	db := postgres.NewConn()

	// create postgres storage
	strg := postgres.NewQuoteStorage(db, logger)

	// create the service layer
	srvc := service.NewQuoteService(logger, strg)

	server := http.NewServer(":8080", logger, srvc).RegisterRoutes()

	server.Run()
}
