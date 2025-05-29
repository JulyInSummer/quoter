package http

import (
	"encoding/json"
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"io"
	"log/slog"
	"net/http"
)

func (s *Server) CreateQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Server.CreateQuote"
	ctx := r.Context()
	s.logger.Info(method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	var req resources.CreateQuoteRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		http_utils.HandleBadRequest(w)
	}

	errors := req.Validate()
	if len(errors) > 0 {
		s.logger.WarnContext(ctx, method+".Validation", slog.Any("errors", errors))
		http_utils.HandleValidationError(w, errors)
		return nil
	}

	quote, err := s.service.CreateQuote(ctx, req.ToDomain())
	if err != nil {
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
