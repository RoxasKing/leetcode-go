package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{[]int{1, 2, 3}},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !compare(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compare(nums1, nums2 [][]int) bool {
	sort.Slice(nums1, func(i, j int) bool {
		var k int
		for nums1[i][k] == nums1[j][k] {
			k++
		}
		return nums1[i][k] < nums1[j][k]
	})
	sort.Slice(nums2, func(i, j int) bool {
		var k int
		for nums2[i][k] == nums2[j][k] {
			k++
		}
		return nums2[i][k] < nums2[j][k]
	})
	return reflect.DeepEqual(nums1, nums2)
}
