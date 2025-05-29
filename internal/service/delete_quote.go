package service

import (
	"context"
	"log/slog"
)

func (q *quoter) DeleteQuote(ctx context.Context, id int) error {
	method := "quoter.DeleteQuote"

	q.logger.DebugContext(ctx, method, slog.Int("quote_id", id))

	err := q.strg.DeleteQuote(ctx, id)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	return nil
}
