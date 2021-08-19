package controller

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/hngi8/task4/config"
	"github.com/hngi8/task4/models"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.html"))
}

// var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Contact(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		Name:     "Abiola Fasanya",
		Email:    "harbiola78@gmail.com",
		Phone:    "2348102307473",
		Location: "Lagos, Nigeria",
	}
	// db := config.DbConn()

	if r.Method != http.MethodPost {
		t, _ := template.ParseFiles("contact.html")
		t.Execute(w, nil)
		return
	}
	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_ = data
	_ = profile

	//insert message database contact
	// newContact := new(models.Contact)
	// newContact.Name = r.FormValue("name")
	// newContact.Email = r.FormValue("email")
	// newContact.Subject = r.FormValue("subject")
	// newContact.Message = r.FormValue("message")

	// insert, err := db.Query("INSERT INTO users (name, email, subject, message) VALUES (newContact.Name, newContact.Email, newContact.Subject, newContact.Message)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
	msg := struct {
		Success bool
		Profile models.Profile
	}{
		Success: true,
		Profile: profile,
	}

	fmt.Println("Data Inserted Successfullly")
	tmpl.Execute(w, msg)
}

func Index(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		Name:     "Abiola Fasanya",
		Email:    "harbiola78@gmail.com",
		Phone:    "2348102307473",
		Location: "Lagos, Nigeria",
	}
	// template, _ := template.ParseFiles("contact.html")
	// template.Execute(w, profile)

	if r.Method != http.MethodPost {
		t, _ := template.ParseFiles("contact.html")
		t.Execute(w, nil)
		return
	}
	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_, _ = data, profile

	msg := struct {
		Success bool
		Profile models.Profile
	}{
		Success: true,
		Profile: profile,
	}

	fmt.Println(msg.Profile, msg.Success)
	fmt.Println("Data Inserted Successfullly")
	tmpl.Execute(w, struct{ Success bool }{true})
}
