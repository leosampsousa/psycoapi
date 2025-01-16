package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/service"
)

type AuthController struct {
	ts *service.TokenService
	us *service.UserService
}

func NewAuthController (ts *service.TokenService, us *service.UserService) *AuthController {
	return &AuthController{
		ts: ts,
		us: us,
	}
}

func (au *AuthController) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var loginDto dto.LoginRequestDto
	if errBind := c.BindJSON(&loginDto); errBind != nil {
		c.IndentedJSON(http.StatusBadRequest, "parâmetros inválidos")
		return
	}

	user, err := au.us.GetUser(ctx, loginDto.Username, loginDto.Password)
	if (err != nil) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"mensagem": "Usuário ou senha inválidos"})
		return
	}

	token, errToken := au.ts.CreateToken(user.Username)
	if (errToken != nil) {
		c.IndentedJSON(errToken.Code, gin.H{"mensagem": errToken.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token" : token})
}

