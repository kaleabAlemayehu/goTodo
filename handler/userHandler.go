package handler

import (
	"context"
	"encoding/json"
	"fmt"
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
			user.getUser(rw, r, int64(id))
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
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 0 {
			id, err := strconv.Atoi(g[0][1])
			if err != nil {
				panic(err)
			}
			user.updateUser(rw, r, int64(id))
		} else {
			http.Error(rw, "Bad Request!", http.StatusBadRequest)
		}
		return
	}
	if r.Method == http.MethodDelete {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 0 {
			id, err := strconv.Atoi(g[0][1])
			if err != nil {
				panic(err)
			}
			user.deleteUser(rw, r, int64(id))
		} else {
			http.Error(rw, "Bad Request!", http.StatusBadRequest)
		}

		return

	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (user *Users) getUser(rw http.ResponseWriter, r *http.Request, id int64) {
	query := db.New(user.Connection)
	encoder := json.NewEncoder(rw)
	u, err := query.GetUser(user.Ctx, id)
	if err != nil {
		http.Error(rw, "USER NOT FOUND", http.StatusNotFound)
		return
	}
	encoder.Encode(u)
}
func (user *Users) getUsers(rw http.ResponseWriter, r *http.Request) {
	query := db.New(user.Connection)
	encoder := json.NewEncoder(rw)
	u, err := query.GetUsers(user.Ctx)
	if err != nil {
		http.Error(rw, "USER NOT FOUND", http.StatusNotFound)
		return
	}
	encoder.Encode(u)

}
func (user *Users) createUser(rw http.ResponseWriter, r *http.Request) {
	var u db.CreateUserParams
	query := db.New(user.Connection)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, "BAD DATA SENT!", http.StatusBadRequest)
		return
	}
	newUser, err := query.CreateUser(user.Ctx, u)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(newUser)
}

func (user *Users) updateUser(rw http.ResponseWriter, r *http.Request, id int64) {
	var u db.UpdateUserParams
	query := db.New(user.Connection)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, "BAD DATA SENT!", http.StatusBadRequest)
		return
	}

	u.ID = id
	err = query.UpdateUser(user.Ctx, u)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (user *Users) deleteUser(rw http.ResponseWriter, r *http.Request, id int64) {

	query := db.New(user.Connection)
	u, err := query.DeleteUser(user.Ctx, id)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(rw)
	err = encoder.Encode(u)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
