package http

import (
	"github.com/JulyInSummer/quoter_app/internal/service"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	serv *http.Server
}

type Handler struct {
	logger  *slog.Logger
	service service.QuoterI
	router  *http.ServeMux
}

func NewServer(address string, handler *Handler) *Server {
	return &Server{
		serv: &http.Server{
			Addr:         address,
			Handler:      handler.router,
			ReadTimeout:  15 * time.Second, // TODO: move to config file
			WriteTimeout: 15 * time.Second, // TODO: move to config file
			IdleTimeout:  60 * time.Second, // TODO: move to config file
		},
	}
}

func (s *Server) Run() error {
	return s.serv.ListenAndServe()
}

func NewHandler(logger *slog.Logger, service service.QuoterI) *Handler {
	h := &Handler{
		logger:  logger,
		service: service,
	}

	handler := http.NewServeMux()

	handler.HandleFunc("POST /quotes", http_utils.Handle(h.CreateQuote))
	handler.HandleFunc("GET /quotes", http_utils.Handle(h.GetAllQuotes))
	handler.HandleFunc("GET /quotes/random", http_utils.Handle(h.GetRandomQuote))
	handler.HandleFunc("DELETE /quotes/{id}", http_utils.Handle(h.DeleteQuote))

	h.router = handler

	return h
}
