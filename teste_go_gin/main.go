package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// projeto represents data about a record projeto.
type pessoa struct {
	ID           string `json:"id"`
	Nome         string `json:"nome"`
	Id_Profissao string `json:"profissao"`
}

type projeto struct {
	ID            string   `json:"id"`
	Nome          string   `json:"nome"`
	Participantes []pessoa `json:"participantes"`
}

// projetos slice to seed record projeto data.
var projetos = []projeto{
	{ID: "1", Nome: "Projeto 1", Participantes: pessoas},
}

var pessoas = []pessoa{
	{ID: "1", Nome: "Bruno", Id_Profissao: "45"},
	{ID: "2", Nome: "Pedro", Id_Profissao: "12"},
	{ID: "3", Nome: "Caio ", Id_Profissao: "13"},
}

func main() {
	router := gin.Default()
	router.GET("/projetos", getProjetos)
	router.GET("/projetos/:id", getProjetoByID)
	router.POST("/projetos", postProjetos)

	router.GET("/pessoas", getPessoas)
	router.GET("/pessoas/:id", getPessoaByID)
	router.POST("/pessoas", postPessoas)

	router.Run("localhost:8080")
}

// getProjetos responds with the list of all projetos as JSON.
func getProjetos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projetos)
}

func getPessoas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pessoas)
}

// postProjetos adds an projeto from JSON received in the request body.
func postProjetos(c *gin.Context) {
	var newProjeto projeto

	// Call BindJSON to bind the received JSON to
	// newProjeto.
	if err := c.BindJSON(&newProjeto); err != nil {
		return
	}

	// Add the new projeto to the slice.
	projetos = append(projetos, newProjeto)
	c.IndentedJSON(http.StatusCreated, newProjeto)
}

func postPessoas(c *gin.Context) {
	var newPessoa pessoa

	// Call BindJSON to bind the received JSON to
	// newProjeto.
	if err := c.BindJSON(&newPessoa); err != nil {
		return
	}

	// Add the new projeto to the slice.
	pessoas = append(pessoas, newPessoa)
	c.IndentedJSON(http.StatusCreated, newPessoa)
}

// getProjetoByID locates the projeto whose ID value matches the id
// parameter sent by the client, then returns that projeto as a response.
func getProjetoByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of projetos, looking for
	// an projeto whose ID value matches the parameter.
	for _, a := range projetos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "projeto not found"})
}

func getPessoaByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of projetos, looking for
	// an projeto whose ID value matches the parameter.
	for _, a := range pessoas {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
}
