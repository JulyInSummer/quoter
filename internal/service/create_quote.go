package service

import (
	"context"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	"log/slog"
)

func (q *quoter) CreateQuote(ctx context.Context, quote domain.Quote) (*domain.Quote, error) {
	method := "quoter.CreateQuote"
	q.logger.DebugContext(ctx, method, slog.Any("quote", quote))

	createdQuote, err := q.strg.CreateQuote(ctx, quote.ToModel())
	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return nil, err
	}

	return &domain.Quote{
		ID:     createdQuote.ID,
		Author: createdQuote.Author,
		Quote:  createdQuote.Quote,
	}, err
}
