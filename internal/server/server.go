package server

import (
	"context"
	"kidstales/internal/server/handlers"
	"kidstales/internal/server/middleware"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	hostPort     = ":80"
	staticDir    = "/static"
	readTimeout  = 15 * time.Second
	writeTimeout = 15 * time.Second
)

type Server struct {
	srv *http.Server
}

func NewServer(ctx context.Context) *Server {
	mx := mux.NewRouter()

	mx.Use(middleware.Logging)

	mx.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))),
	)

	mx.Path("/books").Handler(
		middleware.WithHtmlResponse("list.go.tmpl", handlers.BookList),
	)

	mx.PathPrefix("/book").Handler(
		middleware.WithHtmlResponse("book.go.tmpl", handlers.Book),
	)

	mx.PathPrefix("/proxy").HandlerFunc(handlers.Proxy)

	return &Server{
		srv: &http.Server{
			Addr:         hostPort,
			Handler:      mx,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}
