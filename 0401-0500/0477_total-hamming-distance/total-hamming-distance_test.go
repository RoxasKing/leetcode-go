package main

import "testing"

func Test_totalHammingDistance(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{4, 14, 2}}, 6},
		{"2", args{[]int{4, 14, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalHammingDistance(tt.args.nums); got != tt.want {
				t.Errorf("totalHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
