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

func (s *Service) CreateReceiver(r dtos.CreateReceiverRequest) (*entity.Receiver, error) {
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
	if req.Status == "draft" {
		rcvr, err := entity.NewUpdatebleReceiver(
			req.Id, req.Name, req.Email, req.Doc, req.PixKeyType, req.PixKey, req.Status)
		//TODO make repo verification  if user is indeed draft or active
		if err != nil {
			s.log.Error("invalid user information provided for update", err)
			return err
		}

		_, err = s.repo.UpdateDraft(rcvr)
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

func (s *Service) SearchReceivers(request dtos.SearchRequest) ([]dtos.GetReceiverResponse, error) {
	if request.Limit <= 0 {
		request.Limit = 10
	}
	if request.Limit > 100 {
		request.Limit = 100
	}
	receivers, err := s.repo.Get(request.Query, request.Limit)
	if err != nil {
		s.log.Error(fmt.Sprintf("error finding the receiver with the following query: %s", request.Query), err)
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

func (s *Service) DeleteReceivers(req dtos.DeleReceiverRequest) error {
	return s.repo.Delete(req.Ids...)
}
