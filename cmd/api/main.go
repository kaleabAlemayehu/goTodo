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
	
	collection , context := database.DBConnect()


	r := gin.Default()
	
	r.GET("/", func(ctx *gin.Context) {
		handler.GetAllTodo(ctx, collection, context);
	})
	r.GET("/:id",func(ctx *gin.Context){
		handler.GetTodo(ctx, collection, context)
	})
	r.POST("/", func (ctx *gin.Context)  {
		handler.AddTodo(ctx, collection, context)
	})
	r.PATCH("/:id", func(ctx *gin.Context){
		handler.UpdateTodo(ctx, collection, context)
	})
	r.DELETE("/:id", func(ctx *gin.Context){
		handler.DeleteTodo(ctx, collection, context)
	})

	port := os.Getenv("PORT")
	fmt.Printf("the port is %v:",port)
	r.Run(":" + port)

}