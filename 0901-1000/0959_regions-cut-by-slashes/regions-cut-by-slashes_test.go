package main

import "testing"

func Test_regionsBySlashes(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{" /", "/ "}}, 2},
		{"2", args{[]string{" /", "  "}}, 1},
		{"3", args{[]string{"\\/", "/\\"}}, 4},
		{"4", args{[]string{"/\\", "\\/"}}, 5},
		{"5", args{[]string{"//", "/ "}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regionsBySlashes(tt.args.grid); got != tt.want {
				t.Errorf("regionsBySlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
