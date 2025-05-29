package postgres

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"log/slog"
)

func (q *quoteStorage) GetQuoteByID(ctx context.Context, id int) (*models.Quote, error) {
	method := "quoteStorage.GetQuoteByID"

	q.logger.DebugContext(ctx, method, slog.Int("quote_id", id))

	query := `select id, author, quote from quotes where id = $1`

	var res models.Quote
	err := q.db.QueryRowContext(ctx, query, id).Scan(&res.ID, &res.Author, &res.Quote)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	return &res, nil
}
