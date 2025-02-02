package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/dto"
	"github.com/leosampsousa/psycoapi/internal/service"
	"github.com/leosampsousa/psycoapi/internal/validation"
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
	if errBind := c.BindJSON(&loginDto); errBind != nil || (loginDto.Username == "" || loginDto.Password == "") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "parâmetros inválidos"})
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

func (au *AuthController) Register(c *gin.Context) {
	ctx := c.Request.Context()

	var registerDto dto.RegisterUserDTO
	if errBind := c.BindJSON(&registerDto); errBind != nil || (registerDto.FirstName == "" || registerDto.LastName == "" || registerDto.Password == "" || registerDto.Username == "") {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": "parâmetros inválidos"})
		return
	}

	errInvalidLogin := validation.LoginValidation{}.IsValid(registerDto.Username, registerDto.Password);
	if errInvalidLogin != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": errInvalidLogin.Message})
		return
	}

	err := au.us.CreateUser(ctx, registerDto)
	if (err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"mensagem": err.Message})
		return
	}

	c.Status(http.StatusCreated)
}

func (au *AuthController) IsNewUsername(c *gin.Context) {
	ctx := c.Request.Context()

	var isNewUsernameDto dto.IsNewUsername

	if errBind := c.BindJSON(&isNewUsernameDto); errBind != nil ||  isNewUsernameDto.Username == ""  {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"mensagem": "parâmetros inválidos"})
		return
	}

	print("nome: " + isNewUsernameDto.Username + "\n")

	var isNewUsername = !au.us.AlreadyRegistered(ctx, isNewUsernameDto.Username)

	if(!isNewUsername) {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"mensagem": "esse usuário já existe"})
		return
	}

	c.Status(http.StatusOK)
}
