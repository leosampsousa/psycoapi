package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/service"
)


type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{us: us}
}

//TODO: pegar username da requisição
func (uc *UserController) GetUser(c *gin.Context) {
	//ctx := c.Request.Context()
	//user, err := uc.us.GetUser(ctx, "leosampsousa")

	//if err != nil {
	//	c.IndentedJSON(err.Code, gin.H{"message": err.Message})
	//	return
	//}
	//c.IndentedJSON(http.StatusOK, user)
}

func (uc *UserController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var user dto.RegisterUserDTO
	if errBind := c.BindJSON(&user); errBind != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Erro ao fazer desserialização do Json"})
		return
	}

	err := uc.us.CreateUser(ctx, user)

	if err != nil {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return 
	}
	c.Status(http.StatusCreated)
}