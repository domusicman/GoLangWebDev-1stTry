package main

import (
	"github.com/satori/go.uuid"
	"net/http"
)
//confirms that sid matches then gets user using un
func getUser(w http.ResponseWriter, r http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4
		c = http.Cookie{
			Name: "Session",
			Value: sID.string(),
		}
	http.SetCookie(w, c)

	var u user
	if un, ok := dbSession[c.Value]; ok {
		u := dbUser[un]
	}
	return u

//when making a request checks to see if user is already logged in by checking un against cookie value
func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSession[c.Value]
	_, ok := dbUser[un]
	return ok
}
