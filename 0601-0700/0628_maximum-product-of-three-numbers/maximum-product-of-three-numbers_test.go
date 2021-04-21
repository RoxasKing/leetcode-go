package main

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3}}, 6},
		{"2", args{[]int{1, 2, 3, 4}}, 24},
		{"3", args{[]int{-1, -2, -3}}, -6},
		{"4", args{[]int{-3, -2, 0, 3}}, 18},
		{"5", args{[]int{-3, -2, 1, 3}}, 18},
		{"6", args{[]int{-3, -2, 3, 4}}, 24},
		{"7", args{[]int{-3, -2, -1, 3, 4}}, 24},
		{"8", args{[]int{-3, 2, 3, 4, 5}}, 60},
		{"9", args{[]int{-4, -3, 4, 4, 5}}, 80},
		{"10", args{[]int{-4, -3, 1, 4, 5}}, 60},
		{"11", args{[]int{-4, 0, 0, 0, 5}}, 0},
		{"12", args{[]int{-4, -3, -2, -1}}, -6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
