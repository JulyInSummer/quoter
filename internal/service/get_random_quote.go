package service

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	"log/slog"
)

func (q *quoter) GetRandomQuote(ctx context.Context) (*domain.Quote, error) {
	method := "quoter.GetRandomQuote"

	q.logger.DebugContext(ctx, method)

	quote, err := q.strg.GetRandomQuote(ctx)
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	return &domain.Quote{
		ID:     quote.ID,
		Author: quote.Author,
		Quote:  quote.Quote,
	}, nil
}
