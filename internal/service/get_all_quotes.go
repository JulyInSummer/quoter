package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/service/domain"
	"log/slog"
)

func (q *quoter) GetAllQuotes(ctx context.Context, author string) ([]domain.Quote, error) {
	method := "quoter.GetAllQuotes"

	q.logger.DebugContext(ctx, method)

	res, err := q.strg.GetAllQuotes(ctx, author)

	if err != nil {
		q.logger.ErrorContext(ctx, method, slog.Any("error", err))
		if errors.Is(err, sql.ErrNoRows) {
			return []domain.Quote{}, nil
		}

		return nil, err
	}

	var quotes []domain.Quote
	for _, quote := range res {
		quotes = append(quotes, domain.Quote{
			ID:     quote.ID,
			Author: quote.Author,
			Quote:  quote.Quote,
		})
	}

	return quotes, nil
}
