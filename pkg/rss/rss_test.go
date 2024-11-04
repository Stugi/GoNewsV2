package rss

import (
	"gonews/v2/pkg/storage"
	"reflect"
	"testing"
)

func Test_parseDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Tue, 29 Oct 2024 00:00:00 +0000",
			args: args{
				date: "Tue, 29 Oct 2024 00:00:00 +0000",
			},
			want: 1730160000,
		},
		{
			name: "Tue, 29 Oct 2024 00:00:00 -0700",
			args: args{
				date: "Tue, 29 Oct 2024 00:00:00 -0700",
			},
			want: 1730185200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseDate(tt.args.date); got != tt.want {
				t.Errorf("parseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []storage.Post
		wantErr bool
	}{
		{
			name: "rss",
			args: args{
				url: "https://habr.com/ru/rss/hub/go/all/?fl=ru",
			},
			want:    []storage.Post{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Наверное, можем проверить, что ошибка не возвращается
			_, err := Parse(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewPostFromItem(t *testing.T) {
	type args struct {
		item   Item
		source *storage.Source
	}
	tests := []struct {
		name string
		args args
		want storage.Post
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPostFromItem(tt.args.item, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostFromItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
