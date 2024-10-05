package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/kaleabAlemayehu/goTodo/db"
)

type Users struct {
	Ctx        context.Context
	Connection *pgx.Conn
	Logger     *log.Logger
}

func NewUser(ctx context.Context, con *pgx.Conn, l *log.Logger) *Users {
	return &Users{ctx, con, l}
}

func (user *Users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 0 {
			id, err := strconv.Atoi(g[0][1])
			if err != nil {
				panic(err)
			}
			user.getUser(rw, r, id)
		} else {
			user.getUsers(rw, r)
		}

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

func (user *Users) getUser(rw http.ResponseWriter, r *http.Request, id int) {
	query := db.New(user.Connection)
	encoder := json.NewEncoder(rw)
	u, err := query.GetUser(user.Ctx, int64(id))
	if err != nil {
		http.Error(rw, "USER NOT FOUND", http.StatusNotFound)
	}
	encoder.Encode(u)
}
func (user *Users) getUsers(rw http.ResponseWriter, r *http.Request) {
}
func (user *Users) createUser(rw http.ResponseWriter, r *http.Request) {
}

func (user *Users) updateUser(rw http.ResponseWriter, r *http.Request) {

}
func (user *Users) deleteUser(rw http.ResponseWriter, r *http.Request) {

}
