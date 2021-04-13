package main

import "testing"

func Test_findMagicIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{0, 2, 3, 4, 5}}, 0},
		{"2", args{[]int{1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMagicIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMagicIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
