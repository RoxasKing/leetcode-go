package main

import "testing"

func Test_minStickers(t *testing.T) {
	type args struct {
		stickers []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"with", "example", "science"}, "thehat"}, 3},
		{"2", args{[]string{"notice", "possible"}, "basicbasic"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStickers(tt.args.stickers, tt.args.target); got != tt.want {
				t.Errorf("minStickers() = %v, want %v", got, tt.want)
			}
		})
	}
}
