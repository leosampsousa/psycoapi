package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/leosampsousa/psycoapi/internal/service"
)

type AuthenticationMiddleware struct {
	ts *service.TokenService 
}

func NewAuthenticationMiddleware(ts *service.TokenService) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{ts: ts}
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
            c.JSON(err.Code, err.Message)
            c.Abort()
            return
        }

        c.Set("username", claims["username"])
        c.Next()
    }
}