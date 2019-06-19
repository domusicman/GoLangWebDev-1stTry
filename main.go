package main

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
	// 	Role      string
}

//Establish template var
var tpl *template.Template

var dbUser = map[string]user{}

var dbSession = map[string]string{}

//parse all templates in templates dir and return error if can't
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//handle functions and map to path in url
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", faviconhandler)
	// http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)

}

//first page assign cookie and handle form data
func index(w http.ResponseWriter, r *http.Request) {
	sID, _ := uuid.NewV4()
	fmt.Printf("we got here")

	c := http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(w, &c)

	//handle form
	var u user
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, p, f, l}
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func faviconhandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/watchfavicon.ico")
}
