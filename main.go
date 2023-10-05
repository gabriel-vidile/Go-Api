package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string
	Item      string
	Completed bool
}

var todos = []todo{
	{
		ID:        "1",
		Item:      "Limpar meu quarto",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Comer mingau",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Abraçar o cachorro",
		Completed: false,
	},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {

	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")

}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos/:id", getTodo)

	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run("localhost:9000")
}
