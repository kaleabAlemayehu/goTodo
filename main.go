package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kaleabAlemayehu/goTodo/handler"
	"github.com/kaleabAlemayehu/goTodo/helpers"
)

func main() {
	l := log.New(os.Stdout, "Todo Log >", log.LstdFlags)
	sm := http.NewServeMux()
	ctx, conn := helpers.DBConnect()
	user := handler.NewUser(ctx, conn, l)
	todo := handler.NewTodo(l)
	sm.Handle("/user/", user)
	sm.Handle("/todo/", todo)
	http.ListenAndServe(":9000", sm)
}
