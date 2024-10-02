package handler

import (
	"log"
	"net/http"

	"github.com/kaleabAlemayehu/goTodo/data"
)

type Users struct {
	Logger *log.Logger
}

func NewUser(l *log.Logger) *Users {
	return &Users{l}
}

func (user *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

func (user *Users) getAllUsers(rw http.ResponseWriter, r *http.Request) {
	users := data.GetUsers()

	err := users.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal to json", http.StatusInternalServerError)
	}
}
func (user *Users) createUser(rw http.ResponseWriter, r *http.Request) {

}

func (user *Users) updateUser(rw http.ResponseWriter, r *http.Request) {

}
func (user *Users) deleteUser(rw http.ResponseWriter, r *http.Request) {

}
