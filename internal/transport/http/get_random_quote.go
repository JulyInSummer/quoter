package http

import (
	"database/sql"
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/transport/http/resources"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
)

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Handler.GetRandomQuote"
	ctx := r.Context()

	h.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	quote, err := h.service.GetRandomQuote(ctx)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	resp := resources.QuoteResponse{
		ID:     quote.ID,
		Author: quote.Author,
		Quote:  quote.Quote,
	}

	http_utils.JSON(w, http.StatusOK, resp)
	return nil
}
