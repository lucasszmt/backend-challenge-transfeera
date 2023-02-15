package vo

import "testing"

func TestEmailAddress_GetEmail(t *testing.T) {
	tests := []struct {
		name string
		e    EmailAddress
		want string
	}{
		{"Should return Dwaine Jhonson's EmailKey",
			EmailAddress("dwayneTheRockJhonson@gmail.com"),
			"dwayneTheRockJhonson@gmail.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.GetEmail(); got != tt.want {
				t.Errorf("GetEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmail(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    EmailAddress
		wantErr bool
	}{
		{
			"Should create email with success",
			args{"dwayneTheRockJhonson@gmail.com"},
			EmailAddress("dwayneTheRockJhonson@gmail.com"),
			false,
		},
		{
			"Should not create an email address",
			args{"dwayneTheRockJhonson+gmail.com"},
			EmailAddress(""),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmail(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
