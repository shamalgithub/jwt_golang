package main

import (
	"Go/JWT_Go/initializers"
	
	"net/http"
	"github.com/gin-gonic/gin"
)



func init(){
	initializers.LoadEnvVariables()
	DB := initializers.ConnectToDb()
	initializers.SyncDatabase(DB)
}


func main() {
	
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/todo" , getTodos)
	router.Run("localhost:8080")

}

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