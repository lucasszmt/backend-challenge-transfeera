package receiver

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
	"github.com/lucasszmt/transfeera-challenge/domain/vo"
	"github.com/lucasszmt/transfeera-challenge/infra/log"
	"reflect"
	"testing"
)

type receiverRepoMock struct {
	Err                error
	CreateReceiverMock func(receiver *entity.Receiver) (*entity.Receiver, error)
	UpdateDraftMock    func(receiver *entity.Receiver) (*entity.Receiver, error)
	UpdateValidMock    func(id uuid.UUID, email string) error
	DeleteMock         func(id ...uuid.UUID) error
	GetByIDMock        func(id uuid.UUID) (*dtos.GetReceiverResponse, error)
	GetMock            func(query string, limit int) ([]dtos.GetReceiverResponse, error)
	ListMock           func(page int) ([]dtos.ListReceiversResponse, error)
}

func (r receiverRepoMock) Create(rec *entity.Receiver) (*entity.Receiver, error) {
	switch {
	case r.CreateReceiverMock != nil:
		return r.CreateReceiverMock(rec)
	default:
		return nil, r.Err
	}
}

func (r receiverRepoMock) UpdateDraft(rec *entity.Receiver) (*entity.Receiver, error) {
	switch {
	case r.UpdateDraftMock != nil:
		return r.UpdateDraftMock(rec)
	default:
		return nil, r.Err
	}
}

func (r receiverRepoMock) UpdateValid(id uuid.UUID, email string) error {
	switch {
	case r.UpdateValidMock != nil:
		return r.UpdateValidMock(id, email)
	default:
		return r.Err
	}
}

func (r receiverRepoMock) Delete(id ...uuid.UUID) error {
	switch {
	case r.DeleteMock != nil:
		return r.DeleteMock(id...)
	default:
		return r.Err
	}
}

func (r receiverRepoMock) GetByID(id uuid.UUID) (*dtos.GetReceiverResponse, error) {
	switch {
	case r.GetByIDMock != nil:
		return r.GetByIDMock(id)
	default:
		return nil, r.Err
	}
}

func (r receiverRepoMock) Get(query string, limit int) ([]dtos.GetReceiverResponse, error) {
	switch {
	case r.GetMock != nil:
		return r.GetMock(query, limit)
	default:
		return nil, r.Err
	}
}

func (r receiverRepoMock) List(page int) ([]dtos.ListReceiversResponse, error) {
	switch {
	case r.ListMock != nil:
		return r.ListMock(page)
	default:
		return nil, r.Err
	}
}

func TestService_CreateReceiver(t *testing.T) {
	type fields struct {
		log  log.Logger
		repo Repository
	}
	type args struct {
		r dtos.CreateReceiverRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Should Create an receiver with success",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					CreateReceiverMock: func(receiver *entity.Receiver) (*entity.Receiver, error) {
						return receiver, nil
					},
				},
			},
			args: args{dtos.CreateReceiverRequest{
				Name:       "Anthony Kieds",
				Email:      "rhcp@chilipeppers.com",
				Doc:        "471.550.590-80",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
			}},
			wantErr: false,
		},
		{
			name: "Should return an err for invalid Doc Type",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					CreateReceiverMock: func(receiver *entity.Receiver) (*entity.Receiver, error) {
						return receiver, nil
					},
				},
			},
			args: args{dtos.CreateReceiverRequest{
				Name:       "Anthony Kieds",
				Email:      "rhcp@chilipeppers.com",
				Doc:        "471550580",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
			}},
			wantErr:     true,
			expectedErr: vo.ErrInvalidCPFCNPJ,
		},
		{
			name: "Should get an err on creating a receiver on the repo ",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					Err: sql.ErrConnDone,
				},
			},
			args: args{dtos.CreateReceiverRequest{
				Name:       "Anthony Kieds",
				Email:      "rhcp@chilipeppers.com",
				Doc:        "471.550.590-80",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
			}},
			wantErr:     true,
			expectedErr: sql.ErrConnDone,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:  tt.fields.log,
				repo: tt.fields.repo,
			}
			_, err := s.CreateReceiver(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateReceiver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("CreateReceiver() err = %v, wantErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestService_UpdateReceiver(t *testing.T) {
	type fields struct {
		log  log.Logger
		repo Repository
	}
	type args struct {
		req dtos.UpdateReceiverRequest
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		ExpectedErr error
	}{
		{
			name: "Should Updated Draft User with success",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					UpdateDraftMock: func(r *entity.Receiver) (*entity.Receiver, error) {
						return &entity.Receiver{}, nil
					},
				},
			},
			args: args{dtos.UpdateReceiverRequest{
				Id:         "624b2913-ecf3-4445-9b68-588e41038593",
				Name:       "Chad Smith",
				Email:      "chadsmith@rhcp.com",
				Doc:        "471.550.590-80",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
				Status:     "draft",
			}},
			wantErr:     false,
			ExpectedErr: nil,
		},
		{
			name: "Should Updated Valid User with success",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					UpdateValidMock: func(id uuid.UUID, email string) error {
						return nil
					},
				},
			},
			args: args{dtos.UpdateReceiverRequest{
				Id:     "624b2913-ecf3-4445-9b68-588e41038593",
				Email:  "chadsmith@rhcp.com",
				Status: "valid",
			}},
			wantErr:     false,
			ExpectedErr: nil,
		},
		{
			name: "Should return an err if receiver not found",
			fields: fields{
				log:  log.MockLogger{},
				repo: receiverRepoMock{Err: sql.ErrNoRows},
			},
			args: args{dtos.UpdateReceiverRequest{
				Id:         "624b2913-ecf3-4445-9b68-588e41038593",
				Name:       "Chad Smith",
				Email:      "chadsmith@rhcp.com",
				Doc:        "471.550.590-80",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
				Status:     "draft",
			}},
			wantErr:     true,
			ExpectedErr: sql.ErrNoRows,
		},
		{
			name: "Should return an err if receiver not found",
			fields: fields{
				log:  log.MockLogger{},
				repo: receiverRepoMock{Err: sql.ErrNoRows},
			},
			args: args{dtos.UpdateReceiverRequest{
				Id:         "624b2913-ecf3-4445-9b68-588e41038593",
				Name:       "Chad Smith",
				Email:      "chadsmith@rhcp.com",
				Doc:        "471.550.590-80",
				PixKeyType: vo.CPFKey,
				PixKey:     "471.550.590-80",
				Status:     "active",
			}},
			wantErr:     true,
			ExpectedErr: sql.ErrNoRows,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:  tt.fields.log,
				repo: tt.fields.repo,
			}
			if err := s.UpdateReceiver(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("UpdateReceiver() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_ListReceivers(t *testing.T) {
	type fields struct {
		log  log.Logger
		repo Repository
	}
	type args struct {
		page int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dtos.ListReceiversResponse
		wantErr bool
	}{
		{
			name: "Should return a list of receivers",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{ListMock: func(page int) ([]dtos.ListReceiversResponse, error) {
					return []dtos.ListReceiversResponse{
						{
							Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
							Name:     "Chad Smith",
							Document: "47155059080",
							Status:   "draft",
						}, {
							Id:       uuid.MustParse("65f8b6c5-11a0-4396-a1f2-75cf6f315700"),
							Name:     "Anthony Kieds",
							Document: "47155059089",
							Status:   "active",
						}, {
							Id:       uuid.MustParse("7391dab2-20a3-42be-982e-6f46d6319fad"),
							Name:     "Flea",
							Document: "47155059080",
							Status:   "active",
						},
					}, nil
				}}},
			args: args{3},
			want: []dtos.ListReceiversResponse{
				{
					Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
					Name:     "Chad Smith",
					Document: "47155059080",
					Status:   "draft",
				}, {
					Id:       uuid.MustParse("65f8b6c5-11a0-4396-a1f2-75cf6f315700"),
					Name:     "Anthony Kieds",
					Document: "47155059089",
					Status:   "active",
				}, {
					Id:       uuid.MustParse("7391dab2-20a3-42be-982e-6f46d6319fad"),
					Name:     "Flea",
					Document: "47155059080",
					Status:   "active",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return an error if no receviers found",
			fields: fields{
				log:  log.MockLogger{},
				repo: receiverRepoMock{Err: sql.ErrNoRows},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:  tt.fields.log,
				repo: tt.fields.repo,
			}
			got, err := s.ListReceivers(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListReceivers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListReceivers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_SearchReceivers(t *testing.T) {
	type fields struct {
		log  log.Logger
		repo Repository
	}
	type args struct {
		request dtos.SearchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dtos.GetReceiverResponse
		wantErr bool
	}{
		{
			name: "Should return receivers with a limit of 10",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{
					GetMock: func(query string, limit int) ([]dtos.GetReceiverResponse, error) {
						return []dtos.GetReceiverResponse{
							{
								Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
								Name:     "Jhon Jhones",
								Email:    "jjone@email.ocom",
								Document: "08412535952",
								Pixkey:   "08412535952",
								PixType:  "cpf",
								Status:   "draft",
							},
							{
								Id:       uuid.MustParse("7391dab2-20a3-42be-982e-6f46d6319fad"),
								Name:     "Jhon Cormier",
								Email:    "jcomrier@mail.com",
								Document: "08412535932",
								Pixkey:   "08412535932",
								PixType:  "cpf",
								Status:   "draft",
							},
						}, nil
					},
				},
			},
			args: args{dtos.SearchRequest{
				Query: "Jhon",
			}},
			want: []dtos.GetReceiverResponse{
				{
					Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
					Name:     "Jhon Jhones",
					Email:    "jjone@email.ocom",
					Document: "08412535952",
					Pixkey:   "08412535952",
					PixType:  "cpf",
					Status:   "draft",
				},
				{
					Id:       uuid.MustParse("7391dab2-20a3-42be-982e-6f46d6319fad"),
					Name:     "Jhon Cormier",
					Email:    "jcomrier@mail.com",
					Document: "08412535932",
					Pixkey:   "08412535932",
					PixType:  "cpf",
					Status:   "draft",
				},
			},
			wantErr: false,
		},
		{
			name: "Should return an error on search",
			fields: fields{
				log:  log.MockLogger{},
				repo: receiverRepoMock{Err: sql.ErrNoRows},
			},
			args: args{dtos.SearchRequest{
				Query: "Jhan",
				Limit: 120,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:  tt.fields.log,
				repo: tt.fields.repo,
			}
			got, err := s.SearchReceivers(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchReceivers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchReceivers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetReceiver(t *testing.T) {
	type fields struct {
		log  log.Logger
		repo Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dtos.GetReceiverResponse
		wantErr bool
	}{
		{
			name: "Should get a Receiver with success",
			fields: fields{
				log: log.MockLogger{},
				repo: receiverRepoMock{GetByIDMock: func(id uuid.UUID) (*dtos.GetReceiverResponse, error) {
					return &dtos.GetReceiverResponse{
						Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
						Name:     "Chad Smith",
						Email:    "chadsmith@rhcp.com",
						Document: "471.550.590-80",
						PixType:  "cpf",
						Pixkey:   "471.550.590-80",
						Status:   "draft",
					}, nil
				}},
			},
			args: args{"624b2913-ecf3-4445-9b68-588e41038593"},
			want: &dtos.GetReceiverResponse{
				Id:       uuid.MustParse("624b2913-ecf3-4445-9b68-588e41038593"),
				Name:     "Chad Smith",
				Email:    "chadsmith@rhcp.com",
				Document: "471.550.590-80",
				PixType:  "cpf",
				Pixkey:   "471.550.590-80",
				Status:   "draft",
			},
			wantErr: false,
		},
		{
			name: "Should return an err if receiver not found",
			fields: fields{
				log:  log.MockLogger{},
				repo: receiverRepoMock{Err: ErrReceiverNotFound},
			},
			args:    args{"624b2913-ecf3-4445-9b68-588e41038593"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				log:  tt.fields.log,
				repo: tt.fields.repo,
			}
			got, err := s.GetReceiver(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetReceiver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReceiver() got = %v, want %v", got, tt.want)
			}
		})
	}
}
