package main

import "testing"

func Test_maxLength(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"un", "iq", "ue"}}, 4},
		{"2", args{[]string{"cha", "r", "act", "ers"}}, 6},
		{"3", args{[]string{"abcdefghijklmnopqrstuvwxyz"}}, 26},
		{"4", args{[]string{"yy", "bkhwmpbiisbldzknpm"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLength(tt.args.arr); got != tt.want {
				t.Errorf("maxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
