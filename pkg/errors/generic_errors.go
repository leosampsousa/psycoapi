package errors

import (
	"errors"
	"net/http"
)

//TODO: pensar em uma forma melhor de gerenciar os erros. Melhor separar em erros pra cada serviço
var (
	RecursoNaoEncontrado = errors.New("recurso não encontrado")
	RecursoJaCadastrado = errors.New("recurso ja cadastrado")
	ErroInterno = errors.New("ocorreu um erro interno")
)

func GetHttpStatusFromError(err error) int {
	if err == RecursoNaoEncontrado {
		return http.StatusNotFound
	}

	if(err == RecursoJaCadastrado) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}