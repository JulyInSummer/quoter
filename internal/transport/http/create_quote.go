package http

import (
	"encoding/json"
	"github.com/JulyInSummer/quoter_app/internal/config"
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"io"
	"log/slog"
	"net/http"
)

func (h *Handler) CreateQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Server.CreateQuote"
	ctx := r.Context()
	h.logger.Info(method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	var req resources.CreateQuoteRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		http_utils.HandleBadRequest(w, config.HTTPInvalidBodyMessage)
	}

	errors := req.Validate()
	if len(errors) > 0 {
		h.logger.WarnContext(ctx, method+".Validation", slog.Any("errors", errors))
		http_utils.HandleValidationError(w, errors)
		return nil
	}

	quote, err := h.service.CreateQuote(ctx, req.ToDomain())
	if err != nil {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	resp := resources.QuoteResponse{
		ID:     quote.ID,
		Author: quote.Author,
		Quote:  quote.Quote,
	}

	http_utils.JSON(w, http.StatusCreated, resp)
	return nil
}
