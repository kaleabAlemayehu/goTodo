package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kaleabAlemayehu/goTodo/handler"
)

func main() {
	l := log.New(os.Stdout, "Todo Log >", log.LstdFlags)
	sm := http.NewServeMux()

	user := handler.NewUser(l)
	todo := handler.NewTodo(l)
	sm.Handle("/user/", user)
	sm.Handle("/todo/", todo)
	http.ListenAndServe(":9000", sm)
}