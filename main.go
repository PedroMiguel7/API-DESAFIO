package main

import (
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

// projeto represents data about a record projeto.
type projeto struct {
    ID_Projeto     string  `json:"id"`
    Title  string  `json:"title"`
    Description string  `json:"Description"`
    IDequipe []string `json:"equipe"`
}

type pessoa struct {
    ID_Pessoa     string  `json:"id"`
    Nome  string  `json:"nome"`
    Profissao string  `json:"profissao"`
    ID_Equipe string `json:"equipe"`
    ID_tarefa []string `json:"tarefa"`
}

type equipe struct {
    Nome string `json:"nome"`
    ID_Equipe string   `json:"id"`
}

type tarefa struct {
    ID_Tarefa string `json:"id"`
    Nome string `json:"nome"`
    Description string `json:"description"`
    ID_Project string `json:"ID_Projeto"`
    ID_Equipe string `json:"ID_Equipe"`
    Tempo string `json:"tempo"`
}

// projetos slice to seed record projeto data.
var projetos = []projeto{
    {ID_Projeto: "1", Title: "Central de Relacionamento", Description: "Sugestões", IDequipe: []string{"1", "3"}},
    {ID_Projeto: "2", Title: "Jeru", Description: "talvez de certo", IDequipe: []string{"1", "3"}},
    {ID_Projeto: "3", Title: "Sarah Vaughan and Clifford Brown", Description: "talvez de certo", IDequipe: []string{"1", "3"}},
}

var pessoas = []pessoa{
    {ID_Pessoa: "1", Nome: "Bruno", Profissao: "Dev-Ops", ID_Equipe: "1", ID_tarefa: []string{"1", "4", "5"}},
    {ID_Pessoa: "2", Nome: "Pedro", Profissao: "Back-End", ID_Equipe: "1", ID_tarefa: []string{"1", "3"}},
    {ID_Pessoa: "3", Nome: "Caio",  Profissao: "Front-End", ID_Equipe: "1", ID_tarefa: []string{"3", "2"}},
}

var equipes = []equipe{
    {ID_Equipe: "1", Nome: "Komanda",},
    {ID_Equipe: "2", Nome: "DevsCariri"},
    {ID_Equipe: "3", Nome: "Kariri Inovação"},
}

var tarefas = []tarefa {
    {ID_Tarefa: "1", Nome: "Criação de API REST", Description: "Utilizar GO LANG com Gin", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "2", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "3", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "4", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
    {ID_Tarefa: "5", Nome: "Teste", Description: "Apenas Teste", ID_Project: "1", ID_Equipe: "", Tempo: ""},
}

var menu= []string{
    "Bem vindo!","                                      ","Mais detalhes em: https://github.com/caiosousaf/api_desafio_BrisaNet",
    "Aqui estão todas as rotas disponíveis:",
    "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=",
    "                                      ",
    "-=-=   PROJETOS   =-=-",
    "                                      ",
    "GET:","/projetos","/projetos/:id/tarefas/projetos/:id","/projetos/equipes/:id","/projetos/equipes/:id/members","/projetos/:id/equipes","POST:","/projetos","./projetos/:id/tarefa","PUT:","/projetos/:id","DELETE:","/projetos/:id",
    "                                      ", "------------------------------------","                                      ",
    "-=-=   EQUIPES   =-=-",
    "                                      ", 
    "GET:","/equipes","/equipes/:id","/equipes/member/:id","POST:","/equipes","PUT:","/equipes/:id","DELETE:","/equipes/:id",
    "                                      ","------------------------------------","                                      ",
    "-=-=   MEMBROS   =-=-",
    "                                      ",
    "GET:","/pessoas","/pessoas/:id","/pessoas/:id/tarefas","POST:","/pessoas","PUT:","/pessoas/:id","DELETE:","/pessoas/:id",
    "                                      ", "------------------------------------","                                      ",
    "-=-=   TAREFAS   =-=-",
    "                                      ",
    "GET:","/tarefas","/tarefas/:id","/tarefas/:id/pessoas","POST:","/tarefas","PUT:","/tarefas/:id","DELETE:","/tarefas/:id",
    "                                      ",
    "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=",
}

func main() {
    router := gin.Default()
    router.GET("", gettelainicial)

    router.GET("/projetos", getprojetos)
    router.GET("/projetos/:id", getprojetoByID)
    router.GET("/projetos/:id/tarefas", getTarefasByProject)
    router.GET("/projetos/equipes/:id", getEquipeByID)
    router.GET("/projetos/equipes/:id/members", getMembersInEquipeByID)
    router.GET("/projetos/:id/equipes", getEquipeByProjetobyID)

    router.POST("/projetos", postprojetos)
    router.POST("/projetos/:id/tarefas", postTarefaProjeto)
    router.PUT("/projetos/:id", editProjetoById)
    router.DELETE("/projetos/:id", deleteProjetoById)


    router.GET("/equipes", getEquipes)
    router.GET("/equipes/:id", getEquipeByID)
    router.GET("/equipes/member/:id", getMemberByID)
    router.POST("/equipes", postEquipes)
    router.PUT("/equipes/:id", updateEquipeById)
    router.DELETE("/equipes/:id", deleteEquipeById)
    

    router.GET("/pessoas", getPessoas)
    router.GET("/pessoas/:id", getpessoaByID)
    router.GET("/pessoas/:id/tarefas", getpessoaByIDthetaks)
    router.POST("/pessoas", postpessoas)
    router.PUT("/pessoas/:id", updatePessoaById)
    router.DELETE("/pessoas/:id", deletePessoaById)
    

    router.GET("/tarefas", getTarefas)
    router.GET("/tarefas/:id", getTarefaByID)
    router.GET("/tarefas/:id/pessoas", getTarefaBypeople)
    router.POST("/tarefas", postTarefas)
    router.PUT("/tarefas/:id", editTarefaById)
    router.DELETE("/tarefas/:id", deleteTarefaById)


	//router.Run("localhost:8090")
    port := os.Getenv("PORT")
    router.Run(":"+port)
}

// getprojetos/Pessoas/Equipes responds with the list of all projetos as JSON.
func gettelainicial(c *gin.Context){
    c.IndentedJSON(http.StatusOK, menu)
}

func getprojetos(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, projetos)
}

func getPessoas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pessoas)
}

func getEquipes(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, equipes)
}

func getTarefas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, tarefas)
}

// postprojetos adds an projeto from JSON received in the request body.
func postprojetos(c *gin.Context) {
    var newprojeto projeto

    // Call BindJSON to bind the received JSON to
    // newprojeto.
    if err := c.BindJSON(&newprojeto); err != nil {
        return
    }

    // Add the new projeto to the slice.
    projetos = append(projetos, newprojeto)
    c.IndentedJSON(http.StatusCreated, newprojeto)
}

// getprojetoByID locates the projeto whose ID value matches the id

func getprojetoByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range projetos {
        if a.ID_Projeto == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "projeto not found"})
}

func getEquipeByProjetobyID(c *gin.Context){    
    id := c.Param("id")
    for _, a := range projetos {
        if a.ID_Projeto == id {
            c.IndentedJSON(http.StatusOK, a)
            for _, d := range equipes{
                for _, b:= range a.IDequipe{
                    if d.ID_Equipe == b{
                        c.IndentedJSON(http.StatusOK, d)
                    }
                }
            }
            return
        }
    }
}

    // Delete a project from the list of projects by ID
func deleteProjetoById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range projetos {
        if a.ID_Projeto == id {
            projetos = append(projetos[:i], projetos[i+1:]... )
            return
        }
    }
}

    // Edit a project from the project list by ID
func editProjetoById(c *gin.Context) {
    id := c.Param("id")
    for i := range projetos {
        if projetos[i].ID_Projeto == id {
        c.BindJSON(&projetos[i])
        c.IndentedJSON(http.StatusOK,projetos[i])
        return
        }
    }
}

func postpessoas(c *gin.Context) {
    var newpessoa pessoa

    // Call BindJSON to bind the received JSON to
    // newpessoa.
    if err := c.BindJSON(&newpessoa); err != nil {
        return
    }

    // Add the new pessoa to the slice.
    pessoas = append(pessoas, newpessoa)
    c.IndentedJSON(http.StatusCreated, newpessoa)
}

func postEquipes(c *gin.Context) {
    var newequipe equipe
    // Call BindJSON to bind the received JSON to newpessoa
    if err := c.BindJSON(&newequipe); err != nil {
        return
    }
    // Add the new pessoa to the slice.
    equipes = append(equipes, newequipe)
    c.IndentedJSON(http.StatusCreated, newequipe)
}

func deleteEquipeById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range equipes {
        if a.ID_Equipe == id {
            equipes = append(equipes[:i], equipes[i+1:]... )
            return
        }
    }
}

func updateEquipeById(c *gin.Context) {
    id := c.Param("id")
    for i := range equipes {
        if equipes[i].ID_Equipe == id {
        c.BindJSON(&equipes[i])
        c.IndentedJSON(http.StatusOK,equipes[i])
        return
        }
    }
}

func getpessoaByID(c *gin.Context) {
    id := c.Param("id")

    /* Loop through the list of pessoas, looking for
     an pessoa whose ID value matches the parameter.*/
    for _, a := range pessoas {
        if a.ID_Pessoa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
}

func getMemberByID(c *gin.Context) {
    id := c.Param("id")

    /* Loop through the list of pessoas, looking for
     an pessoa whose ID value matches the parameter.*/
    for _, a := range pessoas {
        if a.ID_Pessoa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Member not found"})
}

func getEquipeByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range equipes {
        if a.ID_Equipe == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Equipe not found"})
}      
// shows the members of a team by id on the screen
func getMembersInEquipeByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range equipes {
		if a.ID_Equipe == id {
			for _, b := range pessoas {
				if b.ID_Equipe == a.ID_Equipe {
					c.IndentedJSON(http.StatusOK, b)
				}
			}
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "equipe not found"})
}

    // Delete a person from the list of people by Id
func deletePessoaById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range pessoas {
        if a.ID_Pessoa == id {
            pessoas = append(pessoas[:i], pessoas[i+1:]... )
            return
        }
    }
}
    // edit a person from a list of people via id
func updatePessoaById(c *gin.Context) {
    id := c.Param("id")
    for i := range pessoas {
        if pessoas[i].ID_Pessoa == id {
        c.BindJSON(&pessoas[i])
        c.IndentedJSON(http.StatusOK,pessoas[i])
        return
        }
    }
}

func getTarefaByID(c *gin.Context) {
    id := c.Param("id")
    for _, a := range tarefas {
        if a.ID_Tarefa == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tarefa not found"})
}

func postTarefas(c *gin.Context) {
    var newtarefa tarefa

    // Call BindJSON to bind the received JSON to
    // newpessoa.
    if err := c.BindJSON(&newtarefa); err != nil {
        return
    }

    // Add the new pessoa to the slice.
    tarefas = append(tarefas, newtarefa)
    c.IndentedJSON(http.StatusCreated, newtarefa)
}

func editTarefaById(c *gin.Context) {
    id := c.Param("id")
    for i := range tarefas {
        if tarefas[i].ID_Tarefa == id {
        c.BindJSON(&tarefas[i])
        c.IndentedJSON(http.StatusOK,tarefas[i])
        return
        }
    }
}
// // Delete a tarefa from the list of tarefas by Id

func deleteTarefaById(c *gin.Context) {
    id := c.Param("id")
    for i, a := range tarefas {
        if a.ID_Tarefa == id {
            tarefas = append(tarefas[:i], tarefas[i+1:]... )
            return
        }
    }
}

func getTarefasByProject(c *gin.Context) {
	id := c.Param("id")
	count := 0
	for _, a := range tarefas {
		if a.ID_Project == id {
			c.IndentedJSON(http.StatusOK, a)
			count+=1
		}
	}
	if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}
// function publish a task in a given project

func postTarefaProjeto(c *gin.Context){
    id := c.Param("id")
	count := 0
    for _, a := range projetos {
        if a.ID_Projeto == id {
                var newtarefa tarefa
                // Call BindJSON to bind the received JSON to new tarefa
                
                if err := c.BindJSON(&newtarefa); err != nil {
                    return
                }
    
                // Add the new Tarefa to the slice by Project.
                tarefas = append(tarefas, newtarefa)
                c.IndentedJSON(http.StatusCreated, newtarefa)
                return
        }  
    }
	if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}

//function that shows people who have a certain task

func getTarefaBypeople(c *gin.Context){
    id := c.Param("id")
	count := 0
	for _, a := range pessoas {
        outrocont := 0
        for range a.ID_tarefa{
		    if a.ID_tarefa[outrocont] == id {
			    c.IndentedJSON(http.StatusOK, a)
			    count+=1
                break
		    }
            outrocont++
        }
	}
	if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
	}
}

//function that shows the tasks of a certain person

func getpessoaByIDthetaks(c *gin.Context){
    id := c.Param("id")
    count := 0
    for _, a := range pessoas {
        if a.ID_Pessoa == id {
            c.IndentedJSON(http.StatusOK, a)
            outrocont := 0
            for range a.ID_tarefa{
                for _, b := range tarefas {
                    if a.ID_tarefa[outrocont] == b.ID_Tarefa{
                        c.IndentedJSON(http.StatusOK, b)
                        outrocont+=1
                    }
                }
            }
            if(outrocont > 0){
                return
            } else{
                c.IndentedJSON(http.StatusNotFound, gin.H{"message": "tarefa not found"})
            }
            count+=1
        }
    }
    if(count > 0){
		return
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pessoa not found"})
	}
}
