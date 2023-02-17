package entity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lucasszmt/transfeera-challenge/domain/vo"
)

type UserStatus int

const (
	Draft UserStatus = iota
	Valid
)

type Receiver struct {
	id     uuid.UUID
	name   vo.Name
	email  vo.EmailAddress
	doc    *vo.CpfCnpj
	pixKey *vo.PixKey
	status UserStatus
}

func NewReceiver(name string, emailAddress string, doc string, pixKeyType vo.PixKeyType,
	pixKey string) (*Receiver, error) {
	r := new(Receiver)
	var err error

	r.id = uuid.New()
	r.name, err = vo.NewName(name)
	if err != nil {
		return nil, err
	}

	if len(emailAddress) > 0 {
		r.email, err = vo.NewEmail(emailAddress)
		if err != nil {
			return nil, err
		}
	}
	r.doc, err = vo.NewCpfCnpj(doc)
	if err != nil {
		return nil, err
	}
	r.pixKey, err = vo.NewPixKey(pixKeyType, pixKey)
	if err != nil {
		return nil, err
	}
	r.status = Draft
	return r, nil
}

// TODO extracts this lot of constructore attributes to structs

func NewUpdatebleReceiver(id string, name string, emailAddress string, doc string, pixKeyType vo.PixKeyType,
	pixKey string) (*Receiver, error) {
	r := new(Receiver)
	var err error

	r.id, err = uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id provided %w", err)
	}
	r.name, err = vo.NewName(name)
	if err != nil {
		return nil, err
	}

	if len(emailAddress) > 0 {
		r.email, err = vo.NewEmail(emailAddress)
		if err != nil {
			return nil, err
		}
	}
	r.doc, err = vo.NewCpfCnpj(doc)
	if err != nil {
		return nil, err
	}
	r.pixKey, err = vo.NewPixKey(pixKeyType, pixKey)
	if err != nil {
		return nil, err
	}
	r.status = Draft
	return r, nil
}

func (r *Receiver) Id() uuid.UUID {
	return r.id
}

func (r *Receiver) Name() string {
	return string(r.name)
}

func (r *Receiver) SetName(name string) error {
	n, err := vo.NewName(name)
	if err != nil {
		return err
	}
	r.name = n
	return nil
}

func (r *Receiver) Email() string {
	return r.email.GetEmail()
}

func (r *Receiver) SetEmail(email string) error {
	e, err := vo.NewEmail(email)
	if err != nil {
		return err
	}
	r.email = e
	return nil
}

func (r *Receiver) Doc() string {
	return r.doc.GetValue()
}

func (r *Receiver) SetDoc(doc *vo.CpfCnpj) {
	r.doc = doc
}

func (r *Receiver) PixKey() *vo.PixKey {
	return r.pixKey
}

func (r *Receiver) SetPixKey(keyType vo.PixKeyType, value string) error {
	pixKey, err := vo.NewPixKey(keyType, value)
	if err != nil {
		return err
	}
	r.pixKey = pixKey
	return nil
}

func (r *Receiver) Status() UserStatus {
	return r.status
}

func (r *Receiver) SetStatus(status UserStatus) {
	r.status = status
}
