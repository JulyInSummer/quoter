package service

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	"github.com/JulyInSummer/quoter_app/internal/storage"
	"log/slog"
)

type QuoterI interface {
	CreateQuote(ctx context.Context, quote domain.Quote) (*domain.Quote, error)
	GetRandomQuote(ctx context.Context) (*domain.Quote, error)
	GetQuoteByID(ctx context.Context, id int) (*domain.Quote, error)
	GetAllQuotes(ctx context.Context, author string) ([]domain.Quote, error)
	DeleteQuote(ctx context.Context, id int) error
}

type quoter struct {
	logger *slog.Logger
	strg   storage.RepoI
}

func NewQuoteService(logger *slog.Logger, strg storage.RepoI) QuoterI {
	return &quoter{
		logger: logger,
		strg:   strg,
	}
}
