package http

import (
	"database/sql"
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
)

func (s *Server) GetRandomQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Server.GetRandomQuote"
	ctx := r.Context()

	s.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	quote, err := s.service.GetRandomQuote(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	resp := resources.QuoteResponse{
		ID:     quote.ID,
		Author: quote.Author,
		Quote:  quote.Quote,
	}

	http_utils.JSON(w, resp)
	return nil
}
