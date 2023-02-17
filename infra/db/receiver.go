package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
)

type Receiver struct {
	db *sqlx.DB
}

func NewReceiver(db *sqlx.DB) *Receiver {
	return &Receiver{db: db}
}

func (r *Receiver) Get(query string) ([]dtos.GetReceiverResponse, error) {
	var resp []dtos.GetReceiverResponse
	if err := r.db.Select(&resp, QueryUser, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, receiver.ErrReceiverNotFound
		}
		return nil, err
	}
	return resp, nil
}

func (r *Receiver) Create(receiver *entity.Receiver) (*entity.Receiver, error) {
	var pixTypeId uuid.UUID
	if err := r.db.Get(&pixTypeId, QueryPixTypeByName, receiver.PixKey().KeyType()); err != nil {
		return receiver, err
	}
	stmt, err := r.db.Preparex(InsertNewReceiverQuery)
	if err != nil {
		return receiver, err
	}
	_, err = stmt.Exec(
		receiver.Id(),
		receiver.Name(),
		receiver.Email(),
		receiver.Doc(),
		receiver.PixKey().Value(),
		pixTypeId.String(),
		receiver.Status())
	if err != nil {
		return receiver, err
	}
	return receiver, nil
}

func (r *Receiver) UpdateDraft(receiver *entity.Receiver) (*entity.Receiver, error) {
	_, err := r.GetByID(receiver.Id())
	if err != nil {
		return nil, err
	}
	var pixTypeId uuid.UUID
	if err := r.db.Get(&pixTypeId, QueryPixTypeByName, receiver.PixKey().KeyType()); err != nil {
		return receiver, err
	}
	stmt, err := r.db.Preparex(UpdateReceiverByID)
	if err != nil {
		return receiver, err
	}
	_, err = stmt.Exec(
		receiver.Name(),
		receiver.Email(),
		receiver.Doc(),
		receiver.PixKey().Value(),
		pixTypeId.String(),
		receiver.Status(),
		receiver.Id())
	return receiver, err
}

func (r *Receiver) UpdateValid(id uuid.UUID, email string) error {
	_, err := r.GetByID(id)
	if err != nil {
		return err
	}
	stmt, err := r.db.Preparex(UpdateReceiverEmailByID)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(email, id); err != nil {
		return err
	}
	return nil
}

func (r *Receiver) GetByID(id uuid.UUID) (*dtos.GetReceiverResponse, error) {
	resp := dtos.GetReceiverResponse{}
	if err := r.db.Get(&resp, QueryUserByID, id.String()); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, receiver.ErrReceiverNotFound
		}
		return nil, err
	}
	return &resp, nil
}

func (r *Receiver) List(page int) ([]dtos.ListReceiversResponse, error) {
	limit := 10
	offset := limit * (page - 1)
	stmt, err := r.db.Preparex(QueryListOfReceivers)
	if err != nil {
		return nil, err
	}
	var resp []dtos.ListReceiversResponse
	err = stmt.Select(&resp, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, receiver.ErrReceiverNotFound
		}
		return nil, err
	}
	return resp, nil
}

func (r *Receiver) Delete(id ...uuid.UUID) error {
	query, args, err := sqlx.In(DelteReceiversByID, id)
	if err != nil {
		return err
	}
	query = r.db.Rebind(query)
	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	fmt.Println(res.RowsAffected())
	return nil
}
