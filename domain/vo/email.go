package vo

import "net/mail"

type EmailAddress string

func NewEmail(address string) (EmailAddress, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", ErrInvalidEmail
	}
	return EmailAddress(addr.Address), err
}

func (e EmailAddress) GetEmail() string {
	return string(e)
}
