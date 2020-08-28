package main

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 2, 3, 1}}, []int{2}},
		{"2", args{[]int{1, 2, 1, 3, 5, 6, 4}}, []int{1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.args.nums)
			mark := make(map[int]bool)
			for _, num := range tt.want {
				mark[num] = true
			}
			if !mark[got] {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
