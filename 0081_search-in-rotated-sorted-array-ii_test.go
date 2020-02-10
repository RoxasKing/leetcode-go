package leetcode

import "testing"

func Test_search0081(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{[]int{1, 2, 1}, 1}, true},
		{"", args{[]int{2, 2, 2, 0, 2, 2}, 0}, true},
		{"", args{[]int{1}, 0}, false},
		{"", args{[]int{1, 1}, 0}, false},
		{"", args{[]int{3, 1, 1}, 0}, false},
		{"", args{[]int{3, 1, 1}, 3}, true},
		{"", args{[]int{1, 1, 3}, 3}, true},
		{"", args{[]int{1, 1, 3, 1}, 3}, true},
		{"", args{[]int{1, 3, 1, 1, 1}, 3}, true},
		{"", args{[]int{0, 1, 1, 2, 0, 0}, 2}, true},
		{"", args{[]int{2, 5, 6, 0, 0, 1, 2}, 0}, true},
		{"", args{[]int{2, 5, 6, 0, 0, 1, 2}, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search0081(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search0081() = %v, want %v", got, tt.want)
			}
		})
	}
}
