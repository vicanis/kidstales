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
	srv               *http.Server
	useTLS            bool
	certPath, keyPath string
}

func NewServer(ctx context.Context) *Server {
	configBuilder := new(config.ServerConfigBuilder).
		WithAddr(config.Env(config.Addr).String()).
		WithTimeout(readTimeout, writeTimeout)

	useTLS := false
	certPath := ""
	keyPath := ""
	if config.Env(config.SSLEnabled).Bool() {
		useTLS = true
		certPath = config.Env(config.SSLCertPath).String()
		keyPath = config.Env(config.SSLKeyPath).String()
		configBuilder = configBuilder.WithSSL(certPath, keyPath)
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

	return &Server{
		srv:      configBuilder.Build(mx),
		useTLS:   useTLS,
		certPath: certPath,
		keyPath:  keyPath,
	}
}

func (s *Server) Start(ctx context.Context) error {
	log.Printf("starting server at %s", s.srv.Addr)

	go func() {
		<-ctx.Done()
		log.Printf("shutdown server")
		_ = s.srv.Close()
	}()

	if s.useTLS {
		return s.srv.ListenAndServeTLS(s.certPath, s.keyPath)
	}

	return s.srv.ListenAndServe()
}
