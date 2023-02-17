package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type receiverServiceMock struct {
	Err                 error
	CreateReceiverMock  func(request dtos.CreateReceiverRequest) (*entity.Receiver, error)
	SearchReceiversMock func(query string) ([]dtos.GetReceiverResponse, error)
	UpdateReceiverMock  func(req dtos.UpdateReceiverRequest) error
	ListReceiversMock   func(page int) ([]dtos.ListReceiversResponse, error)
	GetReceiverMock     func(id string) (*dtos.GetReceiverResponse, error)
}

func (r receiverServiceMock) CreateReceiver(request dtos.CreateReceiverRequest) (*entity.Receiver, error) {
	switch {
	case r.CreateReceiverMock != nil:
		return r.CreateReceiverMock(request)
	default:
		return &entity.Receiver{}, r.Err
	}
}

func (r receiverServiceMock) SearchReceivers(query string) ([]dtos.GetReceiverResponse, error) {
	switch {
	case r.CreateReceiverMock != nil:
		return r.SearchReceiversMock(query)
	default:
		return nil, r.Err
	}
}

func (r receiverServiceMock) UpdateReceiver(req dtos.UpdateReceiverRequest) error {
	switch {
	case r.UpdateReceiverMock != nil:
		return r.UpdateReceiverMock(req)
	default:
		return r.Err
	}
}

func (r receiverServiceMock) ListReceivers(page int) ([]dtos.ListReceiversResponse, error) {
	switch {
	case r.ListReceiversMock != nil:
		return r.ListReceiversMock(page)
	default:
		return nil, r.Err
	}
}

func (r receiverServiceMock) GetReceiver(id string) (*dtos.GetReceiverResponse, error) {
	switch {
	case r.CreateReceiverMock != nil:
		return r.GetReceiverMock(id)
	default:
		return nil, r.Err
	}
}

func Test_receiverHandler_Create(t *testing.T) {
	type args struct {
		service receiver.UseCase
	}
	type expectedResponse struct {
		Code int
		Data interface{}
	}
	const host = "http://localhost"
	const route = "/api/v1/receiver"
	const timeout = int(time.Hour * 1)
	tests := []struct {
		name string
		args args
		req  map[string]interface{}
		want expectedResponse
	}{
		{
			name: "should create user with success",
			args: args{
				receiverServiceMock{
					CreateReceiverMock: func(req dtos.CreateReceiverRequest) (*entity.Receiver, error) {
						return &entity.Receiver{}, nil
					},
				},
			},
			req: map[string]interface{}{
				"name":         "Lucas Szeremeta",
				"email":        "lcuasszmt@gmail.com",
				"doc":          "084.125.359-52",
				"pix_key_type": "cpf",
				"pix_key":      "084.125359-52",
			},
			want: expectedResponse{Code: http.StatusCreated},
		},
		{
			name: "should return a Bad Request status",
			args: args{
				receiverServiceMock{
					CreateReceiverMock: func(req dtos.CreateReceiverRequest) (*entity.Receiver, error) {
						return &entity.Receiver{}, nil
					},
				},
			},
			req:  map[string]interface{}{},
			want: expectedResponse{Code: http.StatusBadRequest},
		},
		{
			name: "should return a Bad Request status for invalid struct(Pix Key required)",
			args: args{
				receiverServiceMock{
					CreateReceiverMock: func(req dtos.CreateReceiverRequest) (*entity.Receiver, error) {
						return &entity.Receiver{}, nil
					},
				},
			},
			req: map[string]interface{}{
				"name":         "Lucas Szeremeta",
				"email":        "lucas@gmail.com",
				"doc":          "08425.359-52",
				"pix_key_type": "cpf",
			},
			want: expectedResponse{Code: http.StatusBadRequest},
		},

		{
			name: "should return a Status Unprocessable Entity ",
			args: args{
				receiverServiceMock{
					CreateReceiverMock: func(req dtos.CreateReceiverRequest) (*entity.Receiver, error) {
						return &entity.Receiver{}, nil
					},
				},
			},
			req: map[string]interface{}{
				"name":         "Lucas Szeremeta",
				"email":        "lucas@gmail.com",
				"doc":          "08425.359-52",
				"pix_key_type": "cpf",
				"pix_key":      "084.125359-52",
			},
			want: expectedResponse{Code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Post(route, NewReceiverHandler(tt.args.service).Create())
			jsonBytes, err := json.Marshal(tt.req)
			require.NoError(t, err)
			jsonBody := bytes.NewReader(jsonBytes)
			req = httptest.NewRequest("POST", fmt.Sprint(host, route), jsonBody)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
		})
	}
}

func Test_receiverHandler_Update(t *testing.T) {
	type args struct {
		service receiver.UseCase
	}
	type expectedResponse struct {
		Code int
		Data interface{}
	}
	const host = "http://localhost"
	const route = "/api/v1/receiver"
	const timeout = int(time.Hour * 1)
	tests := []struct {
		name string
		args args
		req  map[string]interface{}
		want expectedResponse
	}{
		{
			name: "Should update receiver with success",
			args: args{receiverServiceMock{}},
			req: map[string]interface{}{
				"id":           "fbd731d4-d3ac-4305-9d65-72800e821136",
				"name":         "Lucas",
				"email":        "jcena@gmail.com",
				"doc":          "084.125.359-52",
				"pix_key_type": "cpf",
				"pix_key":      "084.125359-52",
				"status":       "draft",
			},
			want: expectedResponse{
				Code: http.StatusOK,
				Data: `{"data":"user with id fbd731d4-d3ac-4305-9d65-72800e821136 updated","status":true}`,
			},
		}, {
			name: "Should return a bad request error response",
			args: args{receiverServiceMock{}},
			req: map[string]interface{}{
				"id":      "fbd731d4-d3ac-4305-9d65-72800e821136",
				"name":    "Lucas",
				"email":   "jcena@gmail.com",
				"pix_key": "084.125359-52",
				"status":  "draft",
			},
			want: expectedResponse{
				Code: http.StatusBadRequest,
				Data: `{"errors":"invalid data request: Key: 'UpdateReceiverRequest.PixKeyType' Error:Field validation for 'PixKeyType' failed on the 'required' tag","status":false}`,
			},
		}, {
			name: "Should return a Status Unprocessable Entity response",
			args: args{receiverServiceMock{
				UpdateReceiverMock: func(req dtos.UpdateReceiverRequest) error {
					return errors.New("error creating pix key, cause : invalid cpf or cnpj provided")
				},
			}},
			req: map[string]interface{}{
				"id":           "fbd731d4-d3ac-4305-9d65-72800e821136",
				"name":         "Lucas",
				"email":        "jcena@gmail.com",
				"doc":          "084.125.359-52",
				"pix_key_type": "cpf",
				"pix_key":      "08435359-52",
				"status":       "valid",
			},
			want: expectedResponse{
				Code: http.StatusUnprocessableEntity,
				Data: `{"errors":"unable to update the receiver requested: error creating pix key, cause : invalid cpf or cnpj provided","status":false}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Patch(route, NewReceiverHandler(tt.args.service).Update())
			jsonBytes, err := json.Marshal(tt.req)
			require.NoError(t, err)
			jsonBody := bytes.NewReader(jsonBytes)
			req = httptest.NewRequest("PATCH", fmt.Sprint(host, route), jsonBody)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expercted [%s] but got [%s]", tt.want, string(body))
		})
	}
}

func Test_receiverHandler_List(t *testing.T) {
	type args struct {
		service receiver.UseCase
	}
	type expectedResponse struct {
		Code int
		Data interface{}
	}
	const host = "http://localhost"
	const route = "/api/v1/receiver"
	const timeout = int(time.Hour * 1)
	tests := []struct {
		name string
		args args
		req  map[string]interface{}
		want expectedResponse
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Patch(route, NewReceiverHandler(tt.args.service).Update())
			jsonBytes, err := json.Marshal(tt.req)
			require.NoError(t, err)
			jsonBody := bytes.NewReader(jsonBytes)
			req = httptest.NewRequest("PATCH", fmt.Sprint(host, route), jsonBody)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expercted [%s] but got [%s]", tt.want, string(body))
		})
	}
}
