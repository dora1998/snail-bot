package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dora1998/snail-bot/utils"
	"testing"
)

func TestTwitterClient_CreateCRCToken(t *testing.T) {
	type fields struct {
		client    *twitter.Client
		envConfig utils.Env
	}
	type args struct {
		crcToken string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "CRC Token (hogefuga)",
			fields: fields{
				client: nil,
				envConfig: utils.Env{
					ConsumerKey:    "",
					ConsumerSecret: "abcdefg",
					Token:          "",
					TokenSecret:    "",
				},
			},
			args: args{
				crcToken: "hogefuga",
			},
			want: "sha256=xbvEo8qITfYS4kMqpaJeOh8i6cp6S8Pleg59MYM+UVM=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TwitterClient{
				client:    tt.fields.client,
				envConfig: tt.fields.envConfig,
			}
			if got := c.CreateCRCToken(tt.args.crcToken); got != tt.want {
				t.Errorf("CreateCRCToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
