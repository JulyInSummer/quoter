package http

import (
	"database/sql"
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
)

func (s *Server) GetAllQuotes(w http.ResponseWriter, r *http.Request) error {
	method := "Server.GetAllQuotes"
	ctx := r.Context()

	s.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	quotes, err := s.service.GetAllQuotes(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	var resp []resources.QuoteResponse
	for _, quote := range quotes {
		resp = append(resp, resources.QuoteResponse{
			ID:     quote.ID,
			Author: quote.Author,
			Quote:  quote.Quote,
		})
	}

	http_utils.JSON(w, resp)
	return nil
}
