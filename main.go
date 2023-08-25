package main

import (
	"net/http"

	"github.com/ArtyomWS/GoWebCourse/controllers"
	"github.com/ArtyomWS/GoWebCourse/templates"
	"github.com/ArtyomWS/GoWebCourse/views"
	"github.com/go-chi/chi/v5"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "text/html")
	http.Error(w, "Page Not Found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.Faq(tpl))

	usersC := controllers.User{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	r.NotFound(notFoundHandler)
	http.ListenAndServe(":3000", r)
}
