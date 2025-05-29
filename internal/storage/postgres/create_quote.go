package postgres

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"log/slog"
)

func (q *quoteStorage) CreateQuote(ctx context.Context, quote models.Quote) (*models.Quote, error) {
	method := "quoteStorage.CreateQuote"

	q.logger.DebugContext(ctx, method, slog.Any("quote", quote))

	query := `insert into quotes (author, quote) values ($1, $2) returning id, author, quote`

	var res models.Quote

	err := q.db.QueryRowContext(ctx, query, quote.Author, quote.Quote).Scan(&res.ID, &res.Author, &res.Quote)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	return &res, nil
}
