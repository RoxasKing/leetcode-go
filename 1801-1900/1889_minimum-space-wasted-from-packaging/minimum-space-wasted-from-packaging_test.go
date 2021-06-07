package main

import "testing"

func Test_minWastedSpace(t *testing.T) {
	type args struct {
		packages []int
		boxes    [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}}, 6},
		{"2", args{[]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}, {3, 4}}}, -1},
		{"3", args{[]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWastedSpace(tt.args.packages, tt.args.boxes); got != tt.want {
				t.Errorf("minWastedSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
