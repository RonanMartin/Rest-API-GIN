package main

// API feita com base em dois tutoriais:
// https://www.youtube.com/watch?v=bj77B59nkTQ
// https://www.youtube.com/watch?v=d_L64KT3SFM

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Minutes referes ao tempo para desenpenhar a tarefa. Mas no original era utilizado para quantidade de livros
// Minutes como quantidade, receberia a função que confere se o livro está disponível e retira 1 (vendido)
type tarefa struct {
	ID        string `json:"id"`
	Item      string `json:"Item"`
	Completed bool   `json:"completed"`
	Minutes   int    `json:"minutes"`
}

var tarefas = []tarefa{
	{ID: "1", Item: "Limpar a Sala", Completed: false, Minutes: 6},
	{ID: "2", Item: "Ler a Bíblia", Completed: false, Minutes: 15},
	{ID: "3", Item: "Estudar GO", Completed: false, Minutes: 120},
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

func retiraMinutos(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem:": "Parâmetro não encontrado"})
		return
	}

	tarefa, err := TarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"mensagem:": "Tarefa não encontrada"})
		return
	}

	if tarefa.Minutes <= 0 {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem:": "Tempo para tarefa acabou, minuto indisponível"})
		return
	}

	tarefa.Minutes -= 1
	context.IndentedJSON(http.StatusOK, tarefa)

}

func adicionaMinutos(context *gin.Context) {
	id, ok := context.GetQuery("id")

	if !ok {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem:": "Parâmetro não encontrado"})
		return
	}

	tarefa, err := TarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"mensagem:": "Tarefa não encontrada"})
		return
	}

	tarefa.Minutes += 1
	context.IndentedJSON(http.StatusOK, tarefa)
}

func getTarefaById(context *gin.Context) {
	id := context.Param("id")
	tarefa, err := TarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Mensagem:": "Tarefa não encontrada"})
	}
	context.IndentedJSON(http.StatusOK, tarefa)
}

func alterarStatusTarefa(context *gin.Context) {
	id := context.Param("id")
	tarefa, err := TarefasById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Mensagem:": "Tarefa não encontrada"})
	}

	tarefa.Completed = !tarefa.Completed

	context.IndentedJSON(http.StatusOK, tarefa)

}

func TarefasById(id string) (*tarefa, error) {
	for i, t := range tarefas {
		if t.ID == id {
			return &tarefas[i], nil
		}
	}
	return nil, errors.New("Tarefa não encontrada")
}

// Para consultar a tarefa o formato será: localhost:9090/tarefas/3 -> sendo este 3 o número do ID
// Para alterar o status do "Completed", o formato é o mesmo que acima, porém não em GET e sim em PATCH.
// Para diminuir 1 minuto o formato será: localhost:9090/retira?id=2 -> para retirar um minuto do id 2
// para dicionar 1 minuto: localhost:9090/adiciona
// o post é feito pelo client e vem em formato JSON
func main() {
	router := gin.Default()
	router.GET("/tarefas", getTarefas)
	router.GET("/tarefas/:id", getTarefaById)
	router.PATCH("/tarefas/:id", alterarStatusTarefa)
	router.PATCH("/retira", retiraMinutos)
	router.PATCH("/adiciona", adicionaMinutos)
	router.POST("/tarefas", addTarefas)
	router.Run("localhost:9090")

}
