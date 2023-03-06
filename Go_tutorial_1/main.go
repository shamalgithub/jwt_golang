package main

import (
	"errors"
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


func todosByQuery( context *gin.Context){
	id , okay := context.GetQuery("id")
	if !okay {
		context.IndentedJSON(http.StatusBadRequest , gin.H{"message":"Missing id query parameter"})

	}else{
		todos , err := getTodosbyId(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message":"Todo not found"})
			return 
		}else{
			context.IndentedJSON(http.StatusOK , todos)
		}
	}

}




func todoById(context *gin.Context){
	id := context.Param("id")
	todoItem , err := getTodosbyId(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"} )
		return 
	}else{
		context.IndentedJSON(http.StatusOK , todoItem)
	}
}

func getTodosbyId(id string)(*todo , error){
	for index , td := range todos {
		if td.ID == id {
			return &todos[index] , nil
		}
	}
	return nil , errors.New("todo item not found")
}


func main(){
	
	router := gin.Default()
	router.GET("/todo" , getTodos)
	router.POST("/post-todos" , addTodos)
	router.GET("/todo/:id" , todoById)
	router.GET("/todos" , todosByQuery)
	router.Run("localhost:8080")
	
	
}