package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_parseDateStr(t *testing.T) {
	type args struct {
		str string
	}
	year := time.Now().Year()

	time1, _ := time.Parse("2006/1/2 MST", fmt.Sprintf("%d/12/31 JST", year))
	time2, _ := time.Parse("2006/1/2 MST", fmt.Sprintf("%d/1/1 JST", year+1))
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "correct pattern",
			args: args{
				str: "12/31",
			},
			want:    time1,
			wantErr: false,
		},
		{
			name: "correct pattern (next year)",
			args: args{
				str: "1/1",
			},
			want:    time2,
			wantErr: false,
		},
		{
			name: "incorrect pattern",
			args: args{
				str: "1234",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "incorrect pattern",
			args: args{
				str: "ほげ",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "incorrect pattern",
			args: args{
				str: "",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDateStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDateStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("parseDateStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}
