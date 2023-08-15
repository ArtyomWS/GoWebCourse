package views

import (
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content Type", "text/html")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("execute template: %v", err)
		http.Error(w, "There is an error executing template", http.StatusInternalServerError)
		return
	}
}
