package main

import "testing"

func Test_escapeGhosts(t *testing.T) {
	type args struct {
		ghosts [][]int
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 0}, {0, 3}}, []int{0, 1}}, true},
		{"2", args{[][]int{{1, 0}}, []int{2, 0}}, false},
		{"3", args{[][]int{{2, 0}}, []int{1, 0}}, false},
		{"4", args{[][]int{{5, 0}, {10, -2}, {0, -5}, {-2, -2}, {-7, 1}}, []int{7, 7}}, false},
		{"5", args{[][]int{{-1, 0}, {0, 1}, {-1, 0}, {0, 1}, {-1, 0}}, []int{0, 0}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := escapeGhosts(tt.args.ghosts, tt.args.target); got != tt.want {
				t.Errorf("escapeGhosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
