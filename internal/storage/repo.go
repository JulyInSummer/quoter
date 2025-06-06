package storage

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
)

type RepoI interface {
	CreateQuote(ctx context.Context, quote models.Quote) (*models.Quote, error)
	GetRandomQuote(ctx context.Context) (*models.Quote, error)
	GetQuoteByID(ctx context.Context, id int) (*models.Quote, error)
	GetAllQuotes(ctx context.Context, author string) ([]models.Quote, error)
	DeleteQuote(ctx context.Context, id int) error
}
