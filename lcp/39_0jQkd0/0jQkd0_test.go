package main

import "testing"

func Test_minimumSwitchingTimes(t *testing.T) {
	type args struct {
		source [][]int
		target [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 3}, {5, 4}}, [][]int{{3, 1}, {6, 5}}}, 1},
		{"2", args{[][]int{{1, 2, 3}, {3, 4, 5}}, [][]int{{1, 3, 5}, {2, 3, 4}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwitchingTimes(tt.args.source, tt.args.target); got != tt.want {
				t.Errorf("minimumSwitchingTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
