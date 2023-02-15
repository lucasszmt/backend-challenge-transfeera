package entity

import "github.com/lucasszmt/transfeera-challenge/domain/vo"

type Receiver struct {
	id     vo.ID
	name   vo.Name
	email  vo.EmailAddress
	doc    *vo.CpfCnpj
	pixKey *vo.PixKey
}

func NewReceiver(
	name string,
	emailAddress string,
	doc string,
	pixKeyType vo.PixKeyType,
	pixKey string) (*Receiver, error) {
	r := new(Receiver)
	var err error

	r.id = vo.NewID()
	r.name, err = vo.NewName(name)
	if err != nil {
		return nil, err
	}
	r.email, err = vo.NewEmail(emailAddress)
	if err != nil {
		return nil, err
	}
	r.doc, err = vo.NewCpfCnpj(doc)
	if err != nil {
		return nil, err
	}
	r.pixKey, err = vo.NewPixKey(pixKeyType, pixKey)
	if err != nil {
		return nil, err
	}
	return r, nil
}
