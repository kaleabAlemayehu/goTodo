package handler

import (
	"log"
	"net/http"
)

type User struct {
	Logger *log.Logger
}

func NewUser(l *log.Logger) *User {
	return &User{l}
}

func (user *User) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user.getAllUsers(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		user.createUser(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		user.updateUser(rw, r)
		return
	}
	if r.Method == http.MethodDelete {
		user.deleteUser(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (user *User) getAllUsers(rw http.ResponseWriter, r *http.Request) {

}
func (user *User) createUser(rw http.ResponseWriter, r *http.Request) {

}

func (user *User) updateUser(rw http.ResponseWriter, r *http.Request) {

}
func (user *User) deleteUser(rw http.ResponseWriter, r *http.Request) {

}
