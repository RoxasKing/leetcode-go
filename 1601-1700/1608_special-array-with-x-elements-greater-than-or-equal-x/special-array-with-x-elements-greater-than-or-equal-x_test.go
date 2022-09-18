package main

import "testing"

func Test_specialArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 5}}, 2},
		{"2", args{[]int{0, 0}}, -1},
		{"3", args{[]int{0, 4, 3, 0, 4}}, 3},
		{"4", args{[]int{3, 6, 7, 7, 0}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specialArray(tt.args.nums); got != tt.want {
				t.Errorf("specialArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
