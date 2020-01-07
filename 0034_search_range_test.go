package My_LeetCode_In_Go

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 2, 3}, 2}, []int{1, 1}},
		{"", args{[]int{1, 3}, 1}, []int{0, 0}},
		{"", args{[]int{5, 7, 7, 8, 8}, 8}, []int{3, 4}},
		{"", args{[]int{1}, 1}, []int{0, 0}},
		{"", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{"", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
