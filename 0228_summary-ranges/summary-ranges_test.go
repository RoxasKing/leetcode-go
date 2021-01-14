package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{[]int{0, 1, 2, 4, 5, 7}}, []string{"0->2", "4->5", "7"}},
		{"2", args{[]int{0, 2, 3, 4, 6, 8, 9}}, []string{"0", "2->4", "6", "8->9"}},
		{"3", args{[]int{}}, []string{}},
		{"4", args{[]int{-1}}, []string{"-1"}},
		{"5", args{[]int{0}}, []string{"0"}},
		{"6", args{[]int{0, 1, 2, 3, 4, 5, 7}}, []string{"0->5", "7"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
