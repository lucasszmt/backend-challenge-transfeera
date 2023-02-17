package receiver

import (
	"fmt"
	"github.com/google/uuid"
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

func (s *Service) CreateReceiver(r dtos.UpdateReceiverRequest) (*entity.Receiver, error) {
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

func (s *Service) UpdateReceiver(req dtos.UpdateReceiverRequest) error {
	rcvr, err := entity.NewUpdatebleReceiver(
		req.Id, req.Name, req.Email, req.Doc, req.PixKeyType, req.PixKey, req.Status)
	if err != nil {
		return err
	}
	if rcvr.Status() == entity.Draft {
		_, err := s.repo.UpdateDraft(rcvr)
		if err != nil {
			return err
		}
		return nil
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return fmt.Errorf("invalid id provided %w", err)
	}
	return s.repo.UpdateValid(id, req.Email)
}

func (s *Service) SearchReceivers(query string) ([]dtos.GetReceiverResponse, error) {
	receivers, err := s.repo.Get(query)
	if err != nil {
		s.log.Error(fmt.Sprintf("error finding the receiver with the following data: %s", query), err)
		return nil, err
	}
	return receivers, nil
}

func (s *Service) GetReceiver(id string) (*dtos.GetReceiverResponse, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id provided: %w", err)
	}

	resp, err := s.repo.GetByID(parsedID)
	if err != nil {
		s.log.Error(fmt.Sprintf("error finding the receiver with the following ID: %s", id), err)
		return nil, err
	}
	return resp, nil
}

func (s *Service) ListReceivers(page int) ([]dtos.ListReceiversResponse, error) {
	list, err := s.repo.List(page)
	if err != nil {
		s.log.Error("error while listing receivers", err)
		return nil, err
	}
	return list, nil
}

func (s *Service) DeleteReceivers(req dtos.DeleReceiverRequester) error {
	return s.repo.Delete(req.Ids...)
}
