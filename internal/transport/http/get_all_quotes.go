package http

import (
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
)

func (s *Server) GetAllQuotes(w http.ResponseWriter, r *http.Request) error {
	method := "Server.GetAllQuotes"
	ctx := r.Context()

	s.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	author := r.URL.Query().Get("author")

	quotes, err := s.service.GetAllQuotes(ctx, author)
	if err != nil {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	resp := make([]resources.QuoteResponse, 0)
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
