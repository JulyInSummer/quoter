package postgres

import (
	"context"
	"log/slog"
)

func (q *quoteStorage) DeleteQuote(ctx context.Context, id int) error {
	method := "quoteStorage.DeleteQuote"
	q.logger.DebugContext(ctx, method, slog.Int("quote_id", id))

	query := `delete from quotes where id = $1`

	_, err := q.db.ExecContext(ctx, query, id)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	return nil
}
