package vo

import (
	"reflect"
	"testing"
)

func TestNewCpfCnpj(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name    string
		args    args
		want    *CpfCnpj
		wantErr bool
	}{
		{
			"Should create a new CPF with success",
			args{"471.550.590-80"},
			&CpfCnpj{"47155059080", CPF},
			false,
		},
		{
			"Should create a new CNPJ with success",
			args{"84.811.956/0001-63"},
			&CpfCnpj{"84811956000163", CNPJ},
			false,
		},
		{
			"Should not get an err while creating a CPF",
			args{"471.550.590-0"},
			nil,
			true,
		},
		{
			"Should get an err while creating a CNPJ",
			args{"84.814.956001-63"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCpfCnpj(tt.args.doc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCpfCnpj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCpfCnpj() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNonDigits(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Should remove non digits from CPF", args{doc: "471.550.590-80"}, "47155059080"},
		{"Should remove dots from CPF", args{doc: "471.550.59080"}, "47155059080"},
		{"Should remove '-' from CPF", args{doc: "471550590-80"}, "47155059080"},
		{"Should remove non digits from CNPJ", args{doc: "84.811.956/0001-63"}, "84811956000163"},
		{"Should remove dots from CNPJ", args{doc: "84.811.956000163"}, "84811956000163"},
		{"Should remove '/' & '-' from CNPJ", args{doc: "84811956/0001-63"}, "84811956000163"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNonDigits(tt.args.doc); got != tt.want {
				t.Errorf("removeNonDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
