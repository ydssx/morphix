package util

import (
	"testing"
)

func TestJsonToMap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{`{"Name": "Alice", "Age": 30, "Address": "123 Main St."}`}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := JsonToMap(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("JsonToMap() = %v, want %v", got, tt.want)
			// }
		})
	}
}
