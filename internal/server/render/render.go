package render

import (
	"kidstales/internal/template"
	"net/http"
	"strings"
)

func RenderTemplate(templateName string, data map[string]any, w http.ResponseWriter) {
	w.Header().Set("content-type", "text/html; charset=utf-8")

	buf := &strings.Builder{}
	template.MustRenderTemplate(templateName, buf, data, false)
	template.MustRenderTemplate("base.go.tmpl", w, buf.String(), true)
}

func RenderError(err error, w http.ResponseWriter) {
	RenderTemplate("error.go.tmpl", map[string]any{
		"ErrorMessage": err.Error(),
	}, w)
}
