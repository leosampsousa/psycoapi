package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/service"
)

type AuthenticationMiddleware struct {
	ts *service.TokenService 
    us *service.UserService
}

func NewAuthenticationMiddleware(ts *service.TokenService, us *service.UserService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{ts: ts, us:us}
}

func (au *AuthenticationMiddleware) ValidateJWT() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"mensagem": "token não enviado"})
            c.Abort()
            return
        }

        tokenParts := strings.Split(tokenString, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"mensagem": "token inválido"})
            c.Abort()
            return
        }

        tokenString = tokenParts[1]

        claims, err := au.ts.VerifyToken(tokenString)
        if err != nil {
            c.JSON(err.Code, gin.H{"mensagem": err.Message})
            c.Abort()
            return
        }

        loggedUser, errGetUser := au.us.GetUser(c.Request.Context(), claims["username"].(string))
        if errGetUser != nil {
            c.JSON(errGetUser.Code, gin.H{"mensagem": errGetUser.Message})
            c.Abort()
            return
        }
        
        c.Set("username", claims["username"])
        c.Set("loggedUser", loggedUser)
        c.Next()
    }
}