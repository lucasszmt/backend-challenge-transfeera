package entity

import (
	"errors"
	vo "github.com/lucasszmt/transfeera-challenge/domain/vo"
	"testing"
)

func TestNewReceiver(t *testing.T) {
	type args struct {
		name         string
		emailAddress string
		doc          string
		pixKeyType   vo.PixKeyType
		pixKey       string
	}
	tests := []struct {
		name    string
		args    args
		want    error
		wantErr bool
	}{
		{
			"Should create a new Receiver with success",
			args{
				name:         "Jhon Cena",
				emailAddress: "jhoncena@gmail.com",
				doc:          "471.550.590-80",
				pixKeyType:   vo.CPFKey,
				pixKey:       "471.550.590-80",
			},
			nil,
			false,
		},
		{
			"Should return an err for invalid name",
			args{
				name:         "Jhon3 Cena",
				emailAddress: "jhoncena@gmail.com",
				doc:          "471.550.590-80",
				pixKeyType:   vo.CPFKey,
				pixKey:       "471.550.590-80",
			},
			vo.ErrInvalidName,
			true,
		},
		{
			"Should return an err for invalid email",
			args{
				name:         "Jhon Cena",
				emailAddress: "jhoncenamail.com",
				doc:          "471.550.590-80",
				pixKeyType:   vo.CPFKey,
				pixKey:       "471.550.590-80",
			},
			vo.ErrInvalidEmail,
			true,
		},
		{
			"Should return an err for invalid CPF",
			args{
				name:         "Jhon Cena",
				emailAddress: "jhoncena@gmail.com",
				doc:          "471.550.59-80",
				pixKeyType:   vo.CPFKey,
				pixKey:       "471.550.590-80",
			},
			vo.ErrInvalidCPFCNPJ,
			true,
		},
		{
			"Should return an error for invalid pixkey",
			args{
				name:         "Jhon Cena",
				emailAddress: "jhoncena@gmail.com",
				doc:          "471.550.590-80",
				pixKeyType:   vo.RandomKey,
				pixKey:       "471.550.590-80",
			},
			vo.ErrInvalidRandomKey,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewReceiver(tt.args.name, tt.args.emailAddress, tt.args.doc, tt.args.pixKeyType, tt.args.pixKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReceiver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !errors.Is(err, tt.want) {
				t.Errorf("NewReceiver() got = %v, want %v", err, tt.want)
			}
		})
	}
}
