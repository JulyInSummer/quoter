package postgres

import (
	"database/sql"
	"github.com/JulyInSummer/quoter_app/internal/storage"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
)

func NewConn() *sql.DB {
	// TODO: move to config
	db, err := sql.Open("postgres", "host=localhost port=5432 user=quoter password=quoter dbname=quoter_db sslmode=disable")
	if err != nil {
		log.Fatalf("failed to open to postgres, err: %v\n", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to connect postgres, err: %v\n", err)
		return nil
	}

	return db
}

type quoteStorage struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewQuoteStorage(db *sql.DB, logger *slog.Logger) storage.RepoI {
	return &quoteStorage{
		db:     db,
		logger: logger,
	}
}
