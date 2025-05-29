package http

import (
	"github.com/JulyInSummer/quoter_app/internal/service"
	http_utils "github.com/JulyInSummer/quoter_app/utils/http"
	"log/slog"
	"net/http"
)

type Server struct {
	serv    *http.Server
	logger  *slog.Logger
	service service.QuoterI
}

func NewServer(address string, logger *slog.Logger, service service.QuoterI) *Server {
	return &Server{
		serv: &http.Server{
			Addr: address,
		},
		logger:  logger,
		service: service,
	}
}

func (s *Server) Run() error {
	return s.serv.ListenAndServe()
}

func (s *Server) RegisterRoutes() *Server {
	handler := http.NewServeMux()

	handler.HandleFunc("POST /quotes", http_utils.Handle(s.CreateQuote))
	handler.HandleFunc("GET /quotes", http_utils.Handle(s.GetAllQuotes))
	handler.HandleFunc("GET /quotes/random", http_utils.Handle(s.GetRandomQuote))
	handler.HandleFunc("DELETE /quotes/{id}", http_utils.Handle(s.DeleteQuote))
	s.serv.Handler = handler
	return s
}
