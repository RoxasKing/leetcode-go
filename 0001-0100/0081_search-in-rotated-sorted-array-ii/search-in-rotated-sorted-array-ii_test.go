package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 1}, 1}, true},
		{"2", args{[]int{2, 2, 2, 0, 2, 2}, 0}, true},
		{"3", args{[]int{1}, 0}, false},
		{"4", args{[]int{1, 1}, 0}, false},
		{"5", args{[]int{3, 1, 1}, 0}, false},
		{"6", args{[]int{3, 1, 1}, 3}, true},
		{"7", args{[]int{1, 1, 3}, 3}, true},
		{"8", args{[]int{1, 1, 3, 1}, 3}, true},
		{"9", args{[]int{1, 3, 1, 1, 1}, 3}, true},
		{"10", args{[]int{0, 1, 1, 2, 0, 0}, 2}, true},
		{"11", args{[]int{2, 5, 6, 0, 0, 1, 2}, 0}, true},
		{"12", args{[]int{2, 5, 6, 0, 0, 1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
