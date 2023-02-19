package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lucasszmt/transfeera-challenge/domain/dtos"
	"github.com/lucasszmt/transfeera-challenge/domain/entity"
	"github.com/lucasszmt/transfeera-challenge/domain/receiver"
	"github.com/lucasszmt/transfeera-challenge/domain/vo"
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
	SearchReceiversMock func(request dtos.SearchRequest) ([]dtos.GetReceiverResponse, error)
	UpdateReceiverMock  func(req dtos.UpdateReceiverRequest) error
	ListReceiversMock   func(page int) ([]dtos.ListReceiversResponse, error)
	GetReceiverMock     func(id string) (*dtos.GetReceiverResponse, error)
	DeleteReceiverMock  func(req dtos.DeleReceiverRequest) error
}

func (r receiverServiceMock) CreateReceiver(request dtos.CreateReceiverRequest) (*entity.Receiver, error) {
	switch {
	case r.CreateReceiverMock != nil:
		return r.CreateReceiverMock(request)
	default:
		return &entity.Receiver{}, r.Err
	}
}

func (r receiverServiceMock) SearchReceivers(request dtos.SearchRequest) ([]dtos.GetReceiverResponse, error) {
	switch {
	case r.SearchReceiversMock != nil:
		return r.SearchReceiversMock(request)
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
	case r.GetReceiverMock != nil:
		return r.GetReceiverMock(id)
	default:
		return nil, r.Err
	}
}

func (r receiverServiceMock) DeleteReceivers(ids dtos.DeleReceiverRequest) error {
	switch {
	case r.DeleteReceiverMock != nil:
		return r.DeleteReceiverMock(ids)
	default:
		return r.Err
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
						return &entity.Receiver{}, vo.ErrInvalidCPFCNPJ
					},
				},
			},
			req: map[string]interface{}{
				"name":         "Lucas Szeremeta",
				"email":        "lucas@gmail.com",
				"doc":          "08sd5.359-52",
				"pix_key_type": "cpf",
				"pix_key":      "084.125359-52",
			},
			want: expectedResponse{Code: http.StatusUnprocessableEntity},
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

			require.True(t, string(body) == tt.want.Data, "invalid response. Expected [%s] but got [%s]", tt.want, string(body))
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
		req  string
		want expectedResponse
	}{
		{
			name: "Should return a list of users with status 200 ok",
			args: args{
				receiverServiceMock{ListReceiversMock: func(page int) ([]dtos.ListReceiversResponse, error) {
					return []dtos.ListReceiversResponse{
						{
							Id:       uuid.MustParse("fbd731d4-d3ac-4305-9d65-72800e821136"),
							Name:     "Lucas",
							Document: "08412535952",
							Status:   "draft",
						},
						{
							Id:       uuid.MustParse("fdd410e8-1bdd-4c8b-9a4b-e0ffd738e38b"),
							Name:     "Lucas Szeremeta",
							Document: "08412535952",
							Status:   "draft",
						},
					}, nil
				}},
			},
			req: "1",
			want: expectedResponse{
				Code: http.StatusOK,
				Data: `{"receivers":[{"id":"fbd731d4-d3ac-4305-9d65-72800e821136","name":"Lucas","document":"08412535952","status":"draft"},{"id":"fdd410e8-1bdd-4c8b-9a4b-e0ffd738e38b","name":"Lucas Szeremeta","document":"08412535952","status":"draft"}],"status":true}`,
			},
		}, {
			name: "Should return a list of users with status 200 ok if no page provided",
			args: args{
				receiverServiceMock{ListReceiversMock: func(page int) ([]dtos.ListReceiversResponse, error) {
					return []dtos.ListReceiversResponse{
						{
							Id:       uuid.MustParse("fbd731d4-d3ac-4305-9d65-72800e821136"),
							Name:     "Lucas",
							Document: "08412535952",
							Status:   "draft",
						},
						{
							Id:       uuid.MustParse("fdd410e8-1bdd-4c8b-9a4b-e0ffd738e38b"),
							Name:     "Lucas Szeremeta",
							Document: "08412535952",
							Status:   "draft",
						},
					}, nil
				}},
			},
			req: "1",
			want: expectedResponse{
				Code: http.StatusOK,
				Data: `{"receivers":[{"id":"fbd731d4-d3ac-4305-9d65-72800e821136","name":"Lucas","document":"08412535952","status":"draft"},{"id":"fdd410e8-1bdd-4c8b-9a4b-e0ffd738e38b","name":"Lucas Szeremeta","document":"08412535952","status":"draft"}],"status":true}`,
			},
		},
		{
			name: "Should return a bad request for a invalid page query param",
			req:  "1a",
			want: expectedResponse{
				Code: http.StatusBadRequest,
				Data: `{"error":"invalid param for pages, it should be an integer","status":true}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Get(route, NewReceiverHandler(tt.args.service).List())

			var param string
			if tt.req != "" {
				param = fmt.Sprint("?page=", tt.req)
			}
			req = httptest.NewRequest("GET", fmt.Sprint(host, route, param), nil)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expected [%s] but got [%s]", tt.want, string(body))
		})
	}
}

func Test_receiverHandler_Get(t *testing.T) {
	type args struct {
		service receiver.UseCase
	}
	type expectedResponse struct {
		Code int
		Data interface{}
	}
	const host = "http://localhost"
	const route = "/api/v1/receiver/:id"
	const timeout = int(time.Hour * 1)
	tests := []struct {
		name string
		args args
		req  string
		want expectedResponse
	}{
		{
			name: "Should Find user with success and return 200",
			args: args{
				receiverServiceMock{GetReceiverMock: func(id string) (*dtos.GetReceiverResponse, error) {
					return &dtos.GetReceiverResponse{
						Id:       uuid.MustParse("9260c278-031f-4d2e-976e-b093dd0452fc"),
						Name:     "Ian Mcgregor",
						Email:    "mcgregor@gmail.com",
						Document: "419.267.660-59",
						Pixkey:   "419.267.660-59",
						PixType:  "cpf",
						Status:   "draft",
					}, nil
				}},
			},
			req: "9260c278-031f-4d2e-976e-b093dd0452fc",
			want: expectedResponse{
				http.StatusOK,
				`{"receivers":{"Id":"9260c278-031f-4d2e-976e-b093dd0452fc","Name":"Ian Mcgregor","Email":"mcgregor@gmail.com","Document":"419.267.660-59","Pixkey":"419.267.660-59","PixType":"cpf","Status":"draft"},"status":true}`,
			},
		},
		{
			name: "Should return a Status not found when receiving a ErrReceiverNotFound",
			args: args{
				receiverServiceMock{GetReceiverMock: func(id string) (*dtos.GetReceiverResponse, error) {
					return nil, receiver.ErrReceiverNotFound
				}},
			},
			req: "9260c278-031f-4d2e-976e-b093dd045xxx",
			want: expectedResponse{
				http.StatusNotFound,
				`{"errors":"receiver not found","status":false}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Get(route, NewReceiverHandler(tt.args.service).Get())

			req = httptest.NewRequest("GET", fmt.Sprint(host, route), nil)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expected [%s] but got [%s]", tt.want, string(body))
		})
	}
}

func Test_receiverHandler_Search(t *testing.T) {
	type args struct {
		service receiver.UseCase
	}
	type expectedResponse struct {
		Code int
		Data interface{}
	}
	const host = "http://localhost"
	const route = "/api/v1/receiver/search"
	const timeout = int(time.Hour * 1)
	tests := []struct {
		name string
		args args
		req  string
		want expectedResponse
	}{
		{
			name: "Should return an error if invalid query params are provided",
			req:  `?limit=a`,
			want: expectedResponse{
				http.StatusBadRequest,
				`{"error":"invalid query params","status":false}`,
			},
		},
		{
			name: "Should fail if no query param is provided",
			want: expectedResponse{
				http.StatusBadRequest,
				`{"error":"query param required","status":false}`,
			},
		},
		{
			name: "Should finds users with success",
			args: args{receiverServiceMock{SearchReceiversMock: func(req dtos.SearchRequest) ([]dtos.GetReceiverResponse, error) {
				return []dtos.GetReceiverResponse{
					{
						Id:       uuid.MustParse("40b0b875-8c6e-456b-99f9-4aea2bcea693"),
						Name:     "Lucas Szeremeta",
						Email:    "lucasszmt@gmail.com",
						Document: "08412535952",
						Pixkey:   "08412535952",
						PixType:  "cpf",
						Status:   "active",
					},
					{
						Id:       uuid.MustParse("450aa274-3824-4076-a6b5-32585b38f900"),
						Name:     "Lucas Szeremeta",
						Email:    "lucasszmt@gmail.com",
						Document: "08412535952",
						Pixkey:   "08412535952",
						PixType:  "cpf",
						Status:   "draft",
					},
				}, nil
			}}},
			req: `?query=08412535952&limit=2`,
			want: expectedResponse{
				http.StatusOK,
				`{"receivers":[{"Id":"40b0b875-8c6e-456b-99f9-4aea2bcea693","Name":"Lucas Szeremeta","Email":"lucasszmt@gmail.com","Document":"08412535952","Pixkey":"08412535952","PixType":"cpf","Status":"active"},{"Id":"450aa274-3824-4076-a6b5-32585b38f900","Name":"Lucas Szeremeta","Email":"lucasszmt@gmail.com","Document":"08412535952","Pixkey":"08412535952","PixType":"cpf","Status":"draft"}],"status":true}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Get(route, NewReceiverHandler(tt.args.service).Search())

			req = httptest.NewRequest("GET", fmt.Sprint(host, route, tt.req), nil)
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expected [%s] but got [%s]", tt.want, string(body))
		})
	}
}

func Test_receiverHandler_Delete(t *testing.T) {
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
			name: "Should return invalid data request",
			args: args{},
			req: map[string]interface{}{
				"ids": "7bc49e64-a623-402c-b9e3-f44b42eda5df",
			},
			want: expectedResponse{
				http.StatusBadRequest,
				`{"errors":"invalid data request","status":false}`,
			},
		},
		{
			name: "Should return invalid data request",
			args: args{},
			req:  map[string]interface{}{},
			want: expectedResponse{
				http.StatusBadRequest,
				`{"errors":"ids field is required, and it needs to be an array","status":false}`,
			},
		},
		{
			name: "Should delete an item with sucess",
			args: args{receiverServiceMock{DeleteReceiverMock: func(req dtos.DeleReceiverRequest) error {
				return nil
			}}},
			req: map[string]interface{}{
				"ids": []string{
					"fbd731d4-d3ac-4305-9d65-72800e821136",
					"7bc49e64-a623-402c-b9e3-f44b42eda5df",
				},
			},
			want: expectedResponse{Code: http.StatusNoContent, Data: ``},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			app := fiber.New()
			app.Delete(route, NewReceiverHandler(tt.args.service).Delete())

			json, err := json.Marshal(tt.req)
			require.NoError(t, err)

			req = httptest.NewRequest("DELETE", fmt.Sprint(host, route), bytes.NewReader(json))
			req.Header.Add("Content-Type", "application/json")
			resp, err := app.Test(req, timeout)
			require.NoErrorf(t, err, "failed to make a test request")
			defer func() {
				resp.Body.Close()
			}()
			require.Equalf(t, tt.want.Code, resp.StatusCode, "failed to ping. Expected status [%d] but got [%d]", tt.want.Code, resp.StatusCode)
			body, err := io.ReadAll(resp.Body)

			require.NoErrorf(t, err, "failed to read response body. Cause  %v", err)

			require.True(t, string(body) == tt.want.Data, "invalid response. Expected [%s] but got [%s]", tt.want, string(body))
		})
	}
}
