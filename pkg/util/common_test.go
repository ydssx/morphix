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

func TestCompareRequests(t *testing.T) {
	type Request struct {
		Field1 string
		Field2 int
	}

	type args struct {
		requests []interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{requests: []interface{}{&Request{Field1: "value", Field2: 10}, &Request{Field1: "value", Field2: 10}}}, true},
		{"", args{requests: []interface{}{&Request{Field1: "value", Field2: 10}, &Request{Field1: "value", Field2: 11}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareRequests(tt.args.requests...); got != tt.want {
				t.Errorf("CompareRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
