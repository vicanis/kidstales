package server

import (
	"context"
	"fmt"
	"io"
	"kidstales/internal/client"
	"kidstales/internal/client/parser"
	"kidstales/internal/server/middleware"
	"kidstales/internal/server/render"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		middleware.WithHtmlResponse(
			"list.go.tmpl",
			func(r *http.Request) (map[string]any, error) {
				client := client.New()

				pageReader, err := client.GetWithCache("/")
				if err != nil {
					return nil, err
				}

				return new(parser.BooksListPageParser).Parse(pageReader)
			},
		),
	)

	mx.PathPrefix("/book").Handler(
		middleware.WithHtmlResponse(
			"book.go.tmpl",
			func(r *http.Request) (map[string]any, error) {
				pageIndex := 0

				pageNumberStr := strings.TrimSpace(r.URL.Query().Get("page"))
				if len(pageNumberStr) > 0 {
					var err error
					pageIndex, err = strconv.Atoi(pageNumberStr)
					if err != nil {
						log.Printf("page number %s parse failed: %v", pageNumberStr, err)
					}
				}

				bookPath := strings.TrimPrefix(r.URL.Path, "/book")

				client := client.New()

				pageReader, err := client.GetWithCache(bookPath)
				if err != nil {
					return nil, err
				}

				data, err := new(parser.BookPageParser).Parse(pageReader)
				if err != nil {
					return nil, err
				}

				if pageIndex > 0 {
					data["PreviousPageURL"] = fmt.Sprintf("%s?page=%d", r.URL.Path, pageIndex-1)
				}

				data["NextPageURL"] = fmt.Sprintf("%s?page=%d", r.URL.Path, pageIndex+1)

				data["CurrentPageNumber"] = pageIndex + 1

				return data, nil
			},
		),
	)

	mx.PathPrefix("/proxy").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			path := strings.TrimPrefix(r.URL.Path, "/proxy")

			client := client.New()

			dataReader, err := client.GetWithCache(path)
			if err != nil {
				render.RenderError(err, w)
				return
			}

			defer dataReader.Close()

			data, err := io.ReadAll(dataReader)
			if err != nil {
				render.RenderError(err, w)
				return
			}

			w.Write(data)
		},
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
