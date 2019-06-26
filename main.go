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
	http.HandleFunc("/bar", bar)
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
	fmt.Printf(c.Value)
	http.SetCookie(w, &c)

	//check to see if user exists and get them
	var u user
	if un, ok := dbSession[c.Value]; ok {
		u = dbUser[un]
	}

	//handle form
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, p, f, l}
		dbSession[c.Value] = un
		dbUser[un] = u

	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)

}

func faviconhandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/watchfavicon.ico")
}

func bar(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	un, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	u := dbUser[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSession[c.Value]
	_, ok := dbUser[un]
	return ok
}

// check if c.cookie matches uname?
// get cookie

// if it doesn't match redirect back to index

//if it does match allow through
