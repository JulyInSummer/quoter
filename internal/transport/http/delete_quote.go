package http

import (
	"errors"
	"github.com/JulyInSummer/quoter_app/internal/config"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Handler.DeleteQuote"
	ctx := r.Context()

	h.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	str := r.PathValue("id")

	if str == "" {
		h.logger.DebugContext(ctx, method, slog.String("id", str))
		http_utils.HandleNotFound(w, config.HTTPNotFoundMessage)
		return nil
	}

	quoteID, err := strconv.Atoi(str)
	if err != nil {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		http_utils.HandleBadRequest(w, config.HTTPInvalidPathParameterMessage)
		return nil
	}

	err = h.service.DeleteQuote(ctx, quoteID)
	if err != nil {
		h.logger.ErrorContext(ctx, method, slog.Any("error", err))
		if errors.Is(err, config.QuoteDoesNotExistError) {
			http_utils.HandleNotFound(w, config.HTTPNotFoundMessage)
			return nil
		}
		return err
	}

	http_utils.JSON(w, http.StatusOK, "Quote was successfully deleted.")
	return nil
}
