package main

import "testing"

func Test_countGoodRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}}, 3},
		{"2", args{[][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodRectangles(tt.args.rectangles); got != tt.want {
				t.Errorf("countGoodRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}
