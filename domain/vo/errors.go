package vo

import "errors"

var (
	ErrInvalidName       = errors.New("invalid name provided")
	ErrInvalidEmail      = errors.New("invalid email provided")
	ErrInvalidCPFCNPJ    = errors.New("invalid cpf or cnpj provided")
	ErrInvalidPhone      = errors.New("invalid phone provided")
	ErrInvalidRandomKey  = errors.New("invalid random key provided")
	ErrInvalidPixKeyType = errors.New("invalid pix key type provided")
)
