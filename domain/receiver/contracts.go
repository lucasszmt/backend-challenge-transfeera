package receiver

import (
	"github.com/google/uuid"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
)

type Writer interface {
	Create(receiver *entity.Receiver) (*entity.Receiver, error)
	UpdateDraft(receiver *entity.Receiver) (*entity.Receiver, error)
	UpdateValid(id uuid.UUID, email string) error
	Delete(id ...uuid.UUID) error
}

type Reader interface {
	GetByID(id uuid.UUID) (*dtos.GetReceiverResponse, error)
	Get(query string) ([]dtos.GetReceiverResponse, error)
	List(page int) ([]dtos.ListReceiversResponse, error)
}

type Repository interface {
	Writer
	Reader
}

type UseCase interface {
	CreateReceiver(request dtos.CreateReceiverRequest) (*entity.Receiver, error)
	SearchReceivers(query string) ([]dtos.GetReceiverResponse, error)
	UpdateReceiver(req dtos.UpdateReceiverRequest) error
	ListReceivers(page int) ([]dtos.ListReceiversResponse, error)
	GetReceiver(id string) (*dtos.GetReceiverResponse, error)
	//DeleteReceivers()
}
