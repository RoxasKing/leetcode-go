package main

import "testing"

func Test_totalFruit(t *testing.T) {
	type args struct {
		tree []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 1}}, 3},
		{"2", args{[]int{0, 1, 2, 2}}, 3},
		{"3", args{[]int{1, 2, 3, 2, 2}}, 4},
		{"4", args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}}, 5},
		{"5", args{[]int{6, 2, 1, 1, 3, 6, 6}}, 3},
		{"6", args{[]int{0, 1, 0, 2}}, 3},
		{"7", args{[]int{1, 0, 1, 4, 1, 4, 1, 2, 3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.tree); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
