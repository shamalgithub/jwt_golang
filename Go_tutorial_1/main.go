package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type todo struct{
	ID string  `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
    {ID: "1", Item: "clean room", Completed: false},
    {ID: "2", Item: "read book", Completed: false},
    {ID: "3", Item: "take a nap", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK , todos)
}


func addTodos(context *gin.Context){
	var newTodo todo

	if err:= context.BindJSON(&newTodo); err!=nil{
		return 
	}

	todos = append(todos , newTodo)
	context.IndentedJSON(http.StatusCreated , newTodo)

}

func main(){
	
	router := gin.Default()
	router.GET("/todo" , getTodos)
	router.POST("/post-todos" , addTodos)
	router.Run("localhost:8080")
	
	
}