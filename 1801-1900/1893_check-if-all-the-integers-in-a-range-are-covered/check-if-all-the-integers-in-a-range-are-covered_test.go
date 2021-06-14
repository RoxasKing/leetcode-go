package main

import "testing"

func Test_isCovered(t *testing.T) {
	type args struct {
		ranges [][]int
		left   int
		right  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 2}, {3, 4}, {5, 6}}, 2, 5}, true},
		{"2", args{[][]int{{1, 10}, {10, 20}}, 21, 21}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCovered(tt.args.ranges, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("isCovered() = %v, want %v", got, tt.want)
			}
		})
	}
}
