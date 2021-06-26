package main

import "testing"

func Test_tupleSameProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 3, 4, 6}}, 8},
		{"2", args{[]int{1, 2, 4, 5, 10}}, 16},
		{"3", args{[]int{2, 3, 4, 6, 8, 12}}, 40},
		{"4", args{[]int{2, 3, 5, 7}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.args.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
