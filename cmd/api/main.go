package main

import (
	"fmt"
	"os"

	"github.com/kaleabAlemayehu/goTodo/api/handler"
	"github.com/kaleabAlemayehu/goTodo/internal/database"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)


func main(){
	err := godotenv.Load(".env")																															
	if err != nil{
		log.Fatal("Error loading .env file from main.go")

	}
	
	database.DBConnect()


	r := gin.Default()
	
	r.GET("/", handler.GetAllTodo)
	r.GET("/:id",handler.GetTodo)
	r.POST("/", handler.AddTodo)
	r.PATCH("/:id", handler.UpdateTodo)
	r.DELETE("/:id", handler.DeleteTodo)

	port := os.Getenv("PORT")
	fmt.Printf("the port is %v:",port)
	r.Run(":" + port)

}