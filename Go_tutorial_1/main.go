package main

import (
	"fmt"
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

// func middlewareFunc1(c *gin.Context) {
// 	fmt.Println("middlewareFunc1 running")
// 	// Next should be used only inside middleware.
// 	//It executes the pending handlers in the chain inside the calling handler.
// 	// example := "hello world"
// 	// return example
// 	c.Next()
// }

// func middlewareFunc2(c *gin.Context) {
// 	fmt.Println("middlewareFunc2 running")
// 	// Abort prevents pending handlers from being called
// 	// c.Abort()
// 	fmt.Println("middlewareFunc2 ending")
// 	c.Next()
// }

func middlewareFunc3() gin.HandlerFunc {
	// run one time logic could inserted here

	
	return func(c *gin.Context) {
		fmt.Println("middlewareFunc3 running")
		c.Next()
	
	}
}

func main(){
	
	router := gin.New()
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	router.Use( middlewareFunc3())

	router.GET("/todo" , getTodos)
	router.POST("/post-todos" , addTodos)
	router.GET("/todo/:id" , todoById)
	router.GET("/todos" , todosByQuery)
	router.Run("localhost:8080")
	
	
}