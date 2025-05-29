package postgres

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"log/slog"
)

func (q *quoteStorage) GetRandomQuote(ctx context.Context) (*models.Quote, error) {
	method := "quoteStorage.GetRandomQuote"

	q.logger.DebugContext(ctx, method)

	query := `select id, author, quote from quotes order by random() limit 1`

	var res models.Quote
	row := q.db.QueryRowContext(ctx, query)

	err := row.Scan(&res.ID, &res.Author, &res.Quote)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	return &res, nil
}
