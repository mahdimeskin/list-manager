package main
import (
	"github.com/gin-gonic/gin"
	"github.com/mahdimeskin/list-manager/models"
	"github.com/mahdimeskin/list-manager/handlers"
)
func main(){
	r := gin.Default()
	models.Connect()

	r.GET("/todos", handlers.GetToDos)
	r.GET("/todos/:id", handlers.GetToDo)
	r.POST("/todos", handlers.CreateToDo)
	r.PATCH("/todos/:id", handlers.UpdateToDo)
	r.DELETE("todos/:id", handlers.DeleteToDo)
	r.Run()
}