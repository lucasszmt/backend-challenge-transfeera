package vo

import "testing"

func TestNewName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Name
		wantErr bool
	}{
		{
			"Should create name with success",
			args{"Jon Cena"},
			Name("Jon Cena"),
			false,
		},
		{
			"Should not create name",
			args{"Jon3 Cena"},
			Name(""),
			true,
		},
		{
			"Should not create last name with a number",
			args{"Jon Cena3"},
			Name(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
