package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hngi8/config"
	"github.com/hngi8/models"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("*.html"))
}

// var tmpl = template.Must(template.ParseFiles("view/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{
		Name:     "Abiola Fasanya",
		Email:    "harbiola78@gmail.com",
		Phone:    "2348102307473",
		Location: "Lagos, Nigeria",
	}

	if r.Method != http.MethodPost {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, profile)
		return
	}

	data := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	_, _, _ = data, profile, tmpl

	db := config.DbConn()

	// insert message database contact
	req := models.Contact{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	insert := "INSERT INTO messages (`name`, `email`, `subject`, `message`) VALUES (?,?,?,?)"

	q, err := db.Query(insert, req.Name, req.Email, req.Subject, req.Message)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Insert Name: %s | Email: %s | Subject: %s | Message: %s", req.Name, req.Email, req.Subject, req.Message)

	defer q.Close()

	msg := struct {
		Success bool
		Profile models.Profile
	}{
		Success: true,
		Profile: profile,
	}
	info := map[string]string{"message": req.Name + " Thank You for Your Message it has been received successfully!"}
	json.NewEncoder(w).Encode(info)

	fmt.Println(msg.Profile, msg.Success)
	fmt.Println("Data Inserted Successfullly")
	// tmpl.ExecuteTemplate(w, "index.html", msg)
	// tmpl.Execute(w, struct{ Success bool }{true})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ContactTable(w http.ResponseWriter, r *http.Request) {
	db := config.DbConn()
	table := "CREATE TABLE IF NOT EXISTS contact (id INTEGER PRIMARY KEY, name VARCHAR(100), email VARCHAR(100), subject VARCHAR(100), message VARCHAR(255))"
	d ,err := db.Query(table)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err.Error())
	}
	_= d
	fmt.Println("Table created successfully")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}