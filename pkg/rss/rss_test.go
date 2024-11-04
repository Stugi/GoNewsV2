package rss

import "testing"

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
