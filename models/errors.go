package models

import "errors"

// ValidationError ...
type ValidationError error

var (
	errorUsername      = ValidationError(errors.New("El username no debe de estar vacio"))
	errorShortUsername = ValidationError(errors.New("El username es demacionado corto"))
	errorLargeUsername = ValidationError(errors.New("El username es demaciado largo"))

	errorEmail = ValidationError(errors.New("Formato invalido de Email"))

	errorPasswordEncryption = ValidationError(errors.New("No es posible cifrar el texto"))

	errorLogin = ValidationError(errors.New("Usuario o password incorrectos"))
)
