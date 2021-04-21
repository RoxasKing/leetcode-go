package main

import "testing"

func Test_minSwapsCouples(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 2, 1, 3}}, 1},
		{"2", args{[]int{3, 2, 0, 1}}, 0},
		{"3", args{[]int{0, 2, 4, 3, 1, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouples(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouples() = %v, want %v", got, tt.want)
			}
		})
	}
}
