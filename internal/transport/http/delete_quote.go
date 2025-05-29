package http

import (
	"github.com/JulyInSummer/quoter_app/internal/config"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
	"strconv"
)

func (s *Server) DeleteQuote(w http.ResponseWriter, r *http.Request) error {
	method := "Server.DeleteQuote"
	ctx := r.Context()

	s.logger.InfoContext(ctx, method, slog.String("method", r.Method), slog.String("url", r.URL.String()))

	str := r.PathValue("id")

	if str == "" {
		s.logger.DebugContext(ctx, method, slog.String("id", str))
		http_utils.HandleNotFound(w, config.HTTPNotFoundMessage)
		return nil
	}

	quoteID, err := strconv.Atoi(str)
	if err != nil {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		http_utils.HandleBadRequest(w, config.HTTPInvalidPathParameterMessage)
		return nil
	}

	err = s.service.DeleteQuote(ctx, quoteID)
	if err != nil {
		s.logger.ErrorContext(ctx, method, slog.Any("error", err))
		return err
	}

	http_utils.JSON(w, http.StatusOK, "Quote was successfully deleted.")
	return nil
}
