package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
	Deleted   bool   `json:"deleted"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
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

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func tooggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")
	index := -1

	for i, t := range todos {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not Found"})
		return
	}

	todos = append(todos[:index], todos[index+1:]...)

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})

}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("todos/:id", getTodo)
	router.PATCH("todos/:id", tooggleTodoStatus)
	router.DELETE("todos/:id", deleteTodo)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
