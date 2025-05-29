package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/config"
	"log/slog"
)

func (q *quoter) DeleteQuote(ctx context.Context, id int) error {
	method := "quoter.DeleteQuote"

	q.logger.DebugContext(ctx, method, slog.Int("quote_id", id))

	_, err := q.strg.GetQuoteByID(ctx, id)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		if errors.Is(err, sql.ErrNoRows) {
			return config.QuoteDoesNotExistError
		}
		return err
	}

	err = q.strg.DeleteQuote(ctx, id)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	return nil
}
