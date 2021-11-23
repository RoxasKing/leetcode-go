package main

import "testing"

func Test_largestComponentSize(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 6, 15, 35}}, 4},
		{"2", args{[]int{20, 50, 9, 63}}, 2},
		{"3", args{[]int{2, 3, 6, 7, 4, 12, 21, 39}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestComponentSize(tt.args.nums); got != tt.want {
				t.Errorf("largestComponentSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
