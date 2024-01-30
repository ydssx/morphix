package llm

import (
	"context"
	"reflect"
	"testing"
)

func TestChat_TextToText(t *testing.T) {
	type args struct {
		ctx    context.Context
		prompt string
	}
	tests := []struct {
		name    string
		c       *Chat
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestChat_TextToText",
			c: NewGoogleAI().StartChat(),
			args: args{
				ctx:    context.Background(),
				prompt: "Hello",
			},
			want:    "Hello",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.TextToText(tt.args.ctx, tt.args.prompt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Chat.TextToText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chat.TextToText() = %v, want %v", got, tt.want)
			}
		})
	}
}
