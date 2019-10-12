package main

import "testing"

func Test_extractStatusIdFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "correct url",
			args:    args{url: "http://twitter.com/d0ra1998/status/1182221777239334912"},
			want:    1182221777239334912,
			wantErr: false,
		},
		{
			name:    "incorrect url",
			args:    args{url: "https://minoru.dev"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractStatusIdFromUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractStatusIdFromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractStatusIdFromUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
