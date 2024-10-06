package template

import (
	"html/template"
	"io"
	"log"
)

var tpl *template.Template

func init() {
	var err error

	tpl, err = template.ParseGlob("/templates/*.go.tmpl")
	if err != nil {
		panic("template parse failed: " + err.Error())
	}

	if len(tpl.Templates()) == 0 {
		panic("no templates were parsed")
	}

	log.Printf("parsed templates count: %d", len(tpl.Templates()))
	for i, eachTemplate := range tpl.Templates() {
		log.Printf("#%d: %s", i+1, eachTemplate.Name())
	}
}

func MustGetTemplate(name string) *template.Template {
	if tpl == nil {
		panic("template was not parsed")
	}

	templateFound := tpl.Lookup(name)

	if templateFound == nil {
		panic("no template found for name " + name)
	}

	return templateFound
}

func MustRenderTemplate(templateName string, w io.Writer, data any, isHtml bool) {
	tpl := MustGetTemplate(templateName)

	var err error
	if isHtml {
		err = tpl.Execute(w, template.HTML(data.(string)))
	} else {
		err = tpl.Execute(w, data)
	}
	if err != nil {
		log.Fatalf("template %s render failed: %v", templateName, err)
	}
}
