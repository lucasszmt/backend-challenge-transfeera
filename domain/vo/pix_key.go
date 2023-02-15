package vo

import (
	"fmt"
	"regexp"
)

var (
	PhoneRegexp     = regexp.MustCompile(`^((?:\+?55)?)([1-9][0-9])(9[0-9]{8})$`)
	RandomKeyRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

type PixKeyType string

const (
	CPFKey    PixKeyType = "cpf"
	CNPJKey   PixKeyType = "cnpj"
	EmailKey  PixKeyType = "email"
	PhoneKey  PixKeyType = "phone"
	RandomKey PixKeyType = "random_key"
)

type PixKey struct {
	value   string
	keyType PixKeyType
}

func NewPixKey(keyType PixKeyType, keyValue string) (*PixKey, error) {
	key := &PixKey{keyType: keyType}
	switch keyType {
	case CPFKey, CNPJKey:
		doc, err := NewCpfCnpj(keyValue)
		if err != nil {
			return nil, fmt.Errorf("error creating pix key, cause : %w", err)
		}
		key.value = doc.GetValue()
	case EmailKey:
		email, err := NewEmail(keyValue)
		if err != nil {
			return nil, fmt.Errorf("error creating pix key, cause : %w", err)
		}
		key.value = email.GetEmail()
	case PhoneKey:
		if err := validatePhone(keyValue); err != nil {
			return nil, fmt.Errorf("error creating pix key, cause : %w", err)
		}
		key.value = keyValue
	case RandomKey:
		if err := validateRandomKey(keyValue); err != nil {
			return nil, fmt.Errorf("error creating pix key, cause : %w", err)
		}
		key.value = keyValue
	default:
		return nil, ErrInvalidPixKeyType
	}
	return key, nil
}

func validatePhone(phoneNumber string) error {
	if !PhoneRegexp.MatchString(phoneNumber) {
		return ErrInvalidPhone
	}
	return nil
}

func validateRandomKey(phoneNumber string) error {
	if !RandomKeyRegexp.MatchString(phoneNumber) {
		return ErrInvalidRandomKey
	}
	return nil
}
