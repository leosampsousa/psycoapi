package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/service"
)


type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{us: us}
}

//TODO: pegar username da requisição
func (ac *UserController) User(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := ac.us.GetUser(ctx, "leosampsousa")

	if err != nil {
		handleError(c, err)
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (ac *UserController) Create(c *gin.Context) {

}

//TODO: extrair isso e criar um handle de erros genericos
func handleError(c *gin.Context, err error) {
	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Nenhum registro encontrado"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "erro interno"})
}