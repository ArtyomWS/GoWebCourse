package controllers

import (
	"fmt"
	"net/http"
)

type User struct {
	Templates struct {
		New Template
	}
}

func (u User) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u User) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Email: ", r.PostForm.Get("email"), "---", r.PostForm.Get("password"))
}
