package main

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}}, true},
		{"2", args{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
