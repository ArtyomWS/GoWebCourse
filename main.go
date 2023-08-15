package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/ArtyomWS/GoWebCourse/controllers"
	"github.com/ArtyomWS/GoWebCourse/views"
	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	t, err := views.Parse(filepath)
	if err != nil {
		log.Printf("parsing template %w", err)
		http.Error(w, "There is an error parsing template", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "text/html")
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func main() {
	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(notFoundHandler)
	http.ListenAndServe(":3000", r)
}
