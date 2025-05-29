package postgres

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/storage/models"
	"log/slog"
)

func (q *quoteStorage) GetAllQuotes(ctx context.Context, author string) ([]models.Quote, error) {
	method := "quoteStorage.GetAllQuotes"

	q.logger.DebugContext(ctx, method)

	query := `select id, author, quote from quotes where $1 = '' or $1 = author order by id desc`

	var res []models.Quote

	rows, err := q.db.QueryContext(ctx, query, author)
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
