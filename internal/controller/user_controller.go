package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/service"
	errHandler "github.com/leosampsousa/psycoapi/pkg/errors"
)


type UserController struct {
	us *service.UserService
}

func NewUserController(us *service.UserService) *UserController {
	return &UserController{us: us}
}

//TODO: pegar username da requisição
func (uc *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, err := uc.us.GetUser(ctx, "leosampsousa")

	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(errHandler.GetHttpStatusFromError(err), gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (uc *UserController) Create(c *gin.Context) {
	ctx := c.Request.Context()

	var user dto.CreateUserDTO
	if errBind := c.BindJSON(&user); errBind != nil {
		c.IndentedJSON(errHandler.GetHttpStatusFromError(errBind), gin.H{"message": errBind.Error()})
		return
	}

	err := uc.us.CreateUser(ctx, user)

	if err != nil {
		c.IndentedJSON(errHandler.GetHttpStatusFromError(err), gin.H{"message": err.Error()})
		return 
	}
	c.Status(http.StatusCreated)
}