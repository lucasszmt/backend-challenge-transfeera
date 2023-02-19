package dtos

import "github.com/google/uuid"

type GetReceiverResponse struct {
	Id       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Document string    `db:"document"`
	Pixkey   string    `db:"pixkey"`
	PixType  string    `db:"pix_type"`
	Status   string    `db:"status"`
}

type ListReceiversResponse struct {
	Id       uuid.UUID `db:"id" json:"id,omitempty"`
	Name     string    `db:"name" json:"name,omitempty"`
	Document string    `db:"document" json:"document,omitempty"`
	Status   string    `db:"status" json:"status,omitempty"`
}
