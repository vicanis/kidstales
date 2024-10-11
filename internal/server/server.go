package server

import (
	"context"
	"kidstales/internal/config"
	"kidstales/internal/server/handlers"
	"kidstales/internal/server/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	staticDir    = "/static"
	readTimeout  = 15 * time.Second
	writeTimeout = 15 * time.Second
)

type Server struct {
	srv *http.Server
}

func NewServer(ctx context.Context) *Server {
	configBuilder := new(config.ServerConfigBuilder).
		WithAddr(config.Env(config.Addr).String()).
		WithTimeout(readTimeout, writeTimeout)

	if config.Env(config.SSLEnabled).Bool() {
		configBuilder = configBuilder.WithSSL(
			config.Env(config.SSLCertPath).String(),
			config.Env(config.SSLKeyPath).String(),
		)
	}

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

	return &Server{srv: configBuilder.Build()}
}

func (s *Server) Start() error {
	log.Printf("starting server at %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}
