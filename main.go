package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tarefa struct {
	ID        string `json:"id"`
	Item      string `json:"Item"`
	Completed bool   `json:"completed"`
}

var tarefas = []tarefa{
	{ID: "1", Item: "Limpar a Sala", Completed: false},
	{ID: "2", Item: "Ler a Bíblia", Completed: false},
	{ID: "3", Item: "Estudar GO", Completed: false},
}

func getTarefas(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tarefas)
}

func addTarefas(context *gin.Context) {
	var newTarefa tarefa

	if err := context.BindJSON(&newTarefa); err != nil {
		return
	}

	tarefas = append(tarefas, newTarefa)

	context.IndentedJSON(http.StatusCreated, newTarefa)
}

func getTarefa(context *gin.Context) {
	id := context.Param("id")
	tarefa, err := getTarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Mensagem:": "Tarefa não encontrada"})
	}
	context.IndentedJSON(http.StatusOK, tarefa)
}

func alterarStatusTarefa(context *gin.Context) {
	id := context.Param("id")
	tarefa, err := getTarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Mensagem:": "Tarefa não encontrada"})
	}

	tarefa.Completed = !tarefa.Completed

	context.IndentedJSON(http.StatusOK, tarefa)

}

func getTarefasById(id string) (*tarefa, error) {
	for i, Bd := range tarefas {
		if Bd.ID == id {
			return &tarefas[i], nil
		}
	}
	return nil, errors.New("Tarefa não encontrada")
}

func main() {
	router := gin.Default()
	router.GET("/tarefas", getTarefas)
	router.GET("/tarefas/:id", getTarefa)
	router.PATCH("/tarefas/:id", alterarStatusTarefa)
	router.POST("/tarefas", addTarefas)
	router.Run("localhost:9090")

}
