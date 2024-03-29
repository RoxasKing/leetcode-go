package main

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 87, 87, 87, 2, 1}}, 13},
		{"2", args{[]int{1, 3, 2, 2, 1}}, 7},
		{"3", args{[]int{1, 3, 2, 2, 2, 1}}, 8},
		{"4", args{[]int{1, 0, 2}}, 5},
		{"5", args{[]int{1, 2, 2}}, 4},
		{"6", args{[]int{1, 2, 1}}, 4},
		{"7", args{[]int{1, 4, 3, 3, 3, 1}}, 8},
		{"8", args{[]int{1, 4, 3, 2, 2, 1}}, 10},
		{"9", args{[]int{1, 2, 3, 4, 5, 3, 2, 1, 2, 6, 5, 4, 3, 3, 2, 1, 1, 3, 3, 3, 4, 2}}, 47},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
