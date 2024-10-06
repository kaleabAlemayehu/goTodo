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

type Todo struct {
	Ctx        context.Context
	Connection *pgx.Conn
	Logger     *log.Logger
}

func NewTodo(ctx context.Context, conn *pgx.Conn, l *log.Logger) *Todo {
	return &Todo{ctx, conn, l}
}

func (todo *Todo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 0 {
			id, err := strconv.Atoi(g[0][1])
			if err != nil {
				todo.Logger.Println(err)
				http.Error(rw, "Parameter is bad!", http.StatusBadRequest)
				return
			}
			todo.getTodo(rw, r, int64(id))
		} else {
			todo.getAllTodos(rw, r)
		}

		return
	}
	if r.Method == http.MethodPost {
		todo.createTodo(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		todo.updateTodo(rw, r)
		return
	}
	if r.Method == http.MethodDelete {
		todo.deleteTodo(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (todo *Todo) getTodo(rw http.ResponseWriter, r *http.Request, id int64) {
}

func (todo *Todo) getAllTodos(rw http.ResponseWriter, r *http.Request) {

}

func (todo *Todo) createTodo(rw http.ResponseWriter, r *http.Request) {
	var t db.CreateTodoParams
	query := db.New(todo.Connection)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		todo.Logger.Println(err)
		http.Error(rw, "Body content is Bad!", http.StatusBadRequest)
		return
	}
	newTodo, err := query.CreateTodo(todo.Ctx, t)
	if err != nil {
		todo.Logger.Println(err)
		http.Error(rw, "Unable to Create Todo", http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(rw)
	encoder.Encode(newTodo)
	return
}
func (todo *Todo) updateTodo(rw http.ResponseWriter, r *http.Request) {

}

func (todo *Todo) deleteTodo(rw http.ResponseWriter, r *http.Request) {

}
