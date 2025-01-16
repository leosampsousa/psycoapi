package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	customError "github.com/leosampsousa/psycoapi/pkg/errors"
)

type TokenService struct {
	Secret []byte
}

func NewTokenService(secret []byte) *TokenService {
	return &TokenService{Secret: secret}
}

func (ts *TokenService) CreateToken(username string) (string, *customError.Error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
        jwt.MapClaims{ 
        "username": username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(), 
        })

    tokenString, err := token.SignedString(ts.Secret)
    if err != nil {
    	return "", customError.NewError(500, "Erro ao criar token") 
    }

 return tokenString, nil
}

func (ts *TokenService) VerifyToken(tokenString string) (jwt.MapClaims, *customError.Error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   return ts.Secret, nil
	})
   
	if err != nil {
	   return nil, customError.NewError(500, "Erro ao verificar token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
   
	return nil, customError.NewError(500, "Token inválido")
 }

