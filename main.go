package main

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
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

var dbUser map[string]user

// var dbSession map[string]string

//parse all templates in templates dir and return error if can't
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", faviconHandler)
	// http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	sID, _ := uuid.NewV4()
	fmt.Printf("we got here")

	c := http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	http.SetCookie(w, &c)

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "images/watchfavicon.ico")
}
