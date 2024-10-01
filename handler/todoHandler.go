package handler

import (
	"log"
	"net/http"
)

type Todo struct {
	Logger *log.Logger
}

func NewTodo(l *log.Logger) *Todo {
	return &Todo{l}
}

func (todo *Todo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		todo.getAllTodos(rw, r)
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

func (todo *Todo) getAllTodos(rw http.ResponseWriter, r *http.Request) {

}

func (todo *Todo) createTodo(rw http.ResponseWriter, r *http.Request) {

}
func (todo *Todo) updateTodo(rw http.ResponseWriter, r *http.Request) {

}

func (todo *Todo) deleteTodo(rw http.ResponseWriter, r *http.Request) {

}
