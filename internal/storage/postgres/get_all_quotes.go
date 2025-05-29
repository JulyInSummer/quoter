package postgres

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"log/slog"
)

func (q *quoteStorage) GetAllQuotes(ctx context.Context) ([]models.Quote, error) {
	method := "quoteStorage.GetAllQuotes"

	q.logger.DebugContext(ctx, method)

	query := `select id, author, quote from quotes order by id desc`

	var res []models.Quote

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var quote models.Quote
		err = rows.Scan(&quote.ID, &quote.Author, &quote.Quote)
		if err != nil {
			q.logger.ErrorContext(ctx, method, slog.Any("error", err))
			return nil, err
		}

		res = append(res, quote)
	}

	return res, nil
}
