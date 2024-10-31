package entity

import (
	"fmt"
	"regexp"
)

func errorParamIsRequired(name string) error {
	return fmt.Errorf("%s é obrigatório", name)
}

func errorParamMinimunValue(name string, length int8) error {
	return fmt.Errorf("%s deve conter no minimo %d caracteres", name, length)
}

func errorParamMaximunValue(name string, length int8) error {
	return fmt.Errorf("%s deve conter no máximo %d caracteres", name, length)
}

func validateEmail(email string) bool {
	var regexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regexEmail.MatchString(email)
}
