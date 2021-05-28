package main

import "testing"

func Test_canChoose(t *testing.T) {
	type args struct {
		groups [][]int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}}, true},
		{"2", args{[][]int{{10, -2}, {1, 2, 3, 4}}, []int{1, 2, 3, 4, 10, -2}}, false},
		{"3", args{[][]int{{1, 2, 3}, {3, 4}}, []int{7, 7, 1, 2, 3, 4, 7, 7}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChoose(tt.args.groups, tt.args.nums); got != tt.want {
				t.Errorf("canChoose() = %v, want %v", got, tt.want)
			}
		})
	}
}
