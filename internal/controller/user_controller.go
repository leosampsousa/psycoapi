package controller

import (
	"net/http"
	"strconv"

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

func (uc *UserController) GetFriends(c *gin.Context) {
	ctx := c.Request.Context()
	loggedUser := c.MustGet("loggedUser").(*dto.UserDTO)

	friends, err := uc.us.GetFriends(ctx, loggedUser.ID)
	if (err != nil) {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"friends": friends})
}

func (uc *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	username := c.Param("username")
	user, err := uc.us.GetUser(ctx, username)
	if (err != nil) {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

func (uc *UserController) AddFriend(c *gin.Context) {
	ctx := c.Request.Context()
	loggedUser := c.MustGet("loggedUser").(*dto.UserDTO)
	friendIdParam := c.Param("friendId")

	friendId, errParse := strconv.ParseInt(friendIdParam, 10, 32)

	if (loggedUser.ID == int32(friendId)) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parametros inválidos"})
		return
	}

	if errParse != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parametros inválidos"})
		return
	}

	err := uc.us.AddFriend(ctx, loggedUser.ID, int32(friendId))
	if (err != nil) {
		c.IndentedJSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.Status(http.StatusOK)
}