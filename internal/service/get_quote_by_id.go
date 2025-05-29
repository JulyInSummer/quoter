package service

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	"log/slog"
)

func (q *quoter) GetQuoteByID(ctx context.Context, id int) (*domain.Quote, error) {
	method := "quoter.GetQuoteByID"

	q.logger.DebugContext(ctx, method, slog.Int("quote_id", id))

	quote, err := q.strg.GetQuoteByID(ctx, id)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	resp := domain.Quote{
		ID:     quote.ID,
		Author: quote.Author,
		Quote:  quote.Quote,
	}

	return &resp, nil
}
