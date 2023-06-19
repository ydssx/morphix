package conf

import "testing"

func TestMustLoad(t *testing.T) {
	type args struct {
		path string
		v    *Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{"../../../configs/gateway", &Config{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustLoad(tt.args.path, tt.args.v)
		})
	}
}
