package controller

import (
	"html/template"
	"net/http"

	"github.com/hngi8/task4/models"
	_ "github.com/go-sql-driver/mysql"
)

var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_ = data
	tmpl.Execute(w, struct{ Success bool }{true})
}