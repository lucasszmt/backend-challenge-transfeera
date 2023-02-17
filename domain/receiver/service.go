package receiver

import (
	"fmt"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
	"github.com/lucasszmt/transfeera-challenge/infra/log"
)

type Service struct {
	log  log.Logger
	repo Repository
}

func NewService(log log.Logger, repo Repository) *Service {
	return &Service{log: log, repo: repo}
}

func (s *Service) CreateReceiver(r dtos.CreateUserRequest) (*entity.Receiver, error) {
	rcv, err := entity.NewReceiver(r.Name, r.Email, r.Doc, r.PixKeyType, r.PixKey)
	if err != nil {
		s.log.Error("error creating the a receiver", err)
		return nil, err
	}
	rcv, err = s.repo.Create(rcv)
	if err != nil {
		s.log.Error("error creating the a receiver", err)
		return nil, err
	}
	return rcv, nil
}

func (s *Service) SearchReceivers(query string) ([]dtos.GetReceiverResponse, error) {
	receivers, err := s.repo.Get(query)
	if err != nil {
		s.log.Error(fmt.Sprintf("error finding the receiver with the following data: %s", query), err)
		return nil, err
	}
	return receivers, nil
}

func (s *Service) ListReceivers(page int) ([]dtos.ListReceiversResponse, error) {
	list, err := s.repo.List(page)
	if err != nil {
		s.log.Error("error while listing receivers", err)
		return nil, err
	}
	return list, nil
}

func (s *Service) DeleteReceivers() {
	//TODO implement me
	panic("implement me")
}
