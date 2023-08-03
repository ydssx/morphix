package jwt

import (
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		uid      int64
		username string
		role     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{1, "ydssx", "admin"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateToken(tt.args.uid, tt.args.username, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Print(token)
			// if got != tt.want {
			// 	t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
