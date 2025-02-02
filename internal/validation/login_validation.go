package validation

import (
	"github.com/leosampsousa/psycoapi/internal/util"
	"github.com/leosampsousa/psycoapi/pkg/errors"
)

type LoginValidation struct {}

func (lv LoginValidation) IsValid(username string, password string) (*errors.Error) {

	if err := IsValidUsername(username); err != nil {
		return err
	}

	if err := IsValidPassword(password); err != nil {
		return err
	}
	
	return nil
}

func IsValidUsername(username string) (*errors.Error) {

	if (util.StringUtils{}.ContainsWhitespace(username)) {
		return &errors.Error{Code: 400, Message: "O nome de usuário não pode conter espaços"}
	}

	if (util.StringUtils{}.HasUpperCase(username)) {
		return &errors.Error{Code: 400, Message: "O nome de usuário não pode conter letras maiúsculas"}
	}

	if (util.StringUtils{}.ContainsSymbols(username)) {
		return &errors.Error{Code: 400, Message: "O numero de usuário nao pode conter simbolos"}
	}

	return nil;
}

func IsValidPassword(password string) (*errors.Error) {

	if (len(password) < 8) {
		return &errors.Error{Code: 400, Message: "A senha deve possuir pelo menos 8 caracteres"}
	}

	if (!util.StringUtils{}.HasLowerCase(password) || !util.StringUtils{}.HasUpperCase(password)) {
		return &errors.Error{Code: 400, Message: "A senha deve conter variação de maiúscula e minúscula"}
	}

	return nil
}