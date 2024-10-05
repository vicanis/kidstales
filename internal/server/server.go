package server

import (
	"context"
	"kidstales/internal/client"
	"kidstales/internal/client/parser"
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

	mx.Path("/parse-main").
		Handler(
			WithHtmlResponse(
				"list.go.tmpl",
				func(r *http.Request) (map[string]any, error) {
					client := client.NewClient()

					pageReader, err := client.MainPage()
					if err != nil {
						return nil, err
					}

					return new(parser.BooksListPageParser).Parse(pageReader)
				},
			),
		)

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
