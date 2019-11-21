package twitter

import (
	"reflect"
	"testing"
)

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
			got, err := ExtractStatusIdFromUrl(tt.args.url)
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

func Test_extractBody(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "correct pattern",
			args: args{
				text: "@assignment_bot 追加 テストタスク 10/16",
			},
			want:    "追加 テストタスク 10/16",
			wantErr: false,
		},
		{
			name: "correct pattern w/comment",
			args: args{
				text: "これは無視 @assignment_bot 削除 テストタスク",
			},
			want:    "削除 テストタスク",
			wantErr: false,
		},
		{
			name: "incorrect pattern",
			args: args{
				text: "@assignment_bot",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractBody(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractBody() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitLongText(t *testing.T) {
	type args struct {
		text      string
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "< maxLength",
			args: args{
				text:      "test",
				maxLength: 10,
			},
			want: []string{"test"},
		},
		{
			name: "> maxLength (all Japanese)",
			args: args{
				text:      "あいうえおかきくけこさしすせそ",
				maxLength: 10,
			},
			want: []string{"あいうえおかきくけこ", "さしすせそ"},
		},
		{
			name: "> maxLength (all English)",
			args: args{
				text:      "abcdefghijklmno",
				maxLength: 10,
			},
			want: []string{"abcdefghij", "klmno"},
		},
		{
			name: "> maxLength (mixed)",
			args: args{
				text:      "abcdefghijアイウエオ🐌",
				maxLength: 10,
			},
			want: []string{"abcdefghij", "アイウエオ🐌"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLongText(tt.args.text, tt.args.maxLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLongText() = %v, want %v", got, tt.want)
			}
		})
	}
}
