package dtos

import (
	"github.com/google/uuid"
	"github.com/lucasszmt/transfeera-challenge/domain/vo"
)

type CreateReceiverRequest struct {
	Name       string        `json:"name,omitempty" validate:"required"`
	Email      string        `json:"email,omitempty" validate:"max=250"`
	Doc        string        `json:"doc,omitempty"`
	PixKeyType vo.PixKeyType `json:"pix_key_type,omitempty" validate:"required"`
	PixKey     string        `json:"pix_key,omitempty" validate:"required,max=140"`
}

type DeleReceiverRequest struct {
	Ids []uuid.UUID `json:"ids" validate:"required"`
}

type UpdateReceiverRequest struct {
	Id         string        `json:"id" validate:"required"`
	Name       string        `json:"name,omitempty" validate:"required"`
	Email      string        `json:"email,omitempty" validate:"max=250"`
	Doc        string        `json:"doc,omitempty"`
	PixKeyType vo.PixKeyType `json:"pix_key_type,omitempty" validate:"required"`
	PixKey     string        `json:"pix_key,omitempty" validate:"required,max=140"`
	Status     string        `json:"status" validate:"required"`
}

type SearchRequest struct {
	Query string `params:"query" validate:"required"`
	Limit int
}
