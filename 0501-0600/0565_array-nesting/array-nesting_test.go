package main

import "testing"

func Test_arrayNesting(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 4, 0, 3, 1, 6, 2}}, 4},
		{"2", args{[]int{0, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayNesting(tt.args.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
