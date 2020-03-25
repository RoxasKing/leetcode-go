package leetcode

import "testing"

func Test_surfaceArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{2}}}, 10},
		{"2", args{[][]int{{1, 2}, {3, 4}}}, 34},
		{"3", args{[][]int{{1, 0}, {0, 2}}}, 16},
		{"4", args{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}}, 32},
		{"5", args{[][]int{{2, 2, 2}, {2, 1, 2}, {2, 2, 2}}}, 46},
		{"6", args{[][]int{{2, 3}, {2, 4}}}, 34},
		{"7", args{[][]int{{1, 4}, {1, 1}}}, 28},
		{"8", args{[][]int{{3, 3, 3}, {3, 4, 5}, {5, 0, 4}}}, 78},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := surfaceArea(tt.args.grid); got != tt.want {
				t.Errorf("surfaceArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
