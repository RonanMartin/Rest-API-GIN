package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Bdado struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var Bdados = []Bdados{
	{ID: "1", Item: "Cleam Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getBDADOS(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Bdados)
}

func main() {
	router := gin.Defaut()
	router.GET("/Bdados", getBDADOS)
	router.Run("localhost:9090")

}
