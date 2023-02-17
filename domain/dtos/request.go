package dtos

import "github.com/lucasszmt/transfeera-challenge/domain/vo"

type CreateUserRequest struct {
	Name       string        `json:"name,omitempty" validate:"required"`
	Email      string        `json:"email,omitempty" validate:"max=250"`
	Doc        string        `json:"doc,omitempty"`
	PixKeyType vo.PixKeyType `json:"pix_key_type,omitempty" validate:"required"`
	PixKey     string        `json:"pix_key,omitempty" validate:"required,max=140"`
}
