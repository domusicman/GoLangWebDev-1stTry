package main

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	// 	Password  string
	// 	FirstName string
	// 	LastName  string
	// 	Role      string
}

//Establish template var
var tpl *template.Template

// var dbUser map[string]user
// var dbSession map[string]string

//parse all templates in templates dir and return error if can't
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	// http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	fmt.Printf("we got here")

	c := http.Cookie{
		Name:  "session",
		Value: "value",
	}
	http.SetCookie(w, &c)

}