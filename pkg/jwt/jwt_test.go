package jwt

import "testing"

func TestGenerateToken(t *testing.T) {
	type args struct {
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
		{"",args{"ydssx","admin"},"",false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenerateToken(tt.args.username, tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if got != tt.want {
			// 	t.Errorf("GenerateToken() = %v, want %v", got, tt.want)
			// }
		})
	}
}
