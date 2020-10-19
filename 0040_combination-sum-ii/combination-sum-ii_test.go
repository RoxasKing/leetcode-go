package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{2, 5, 2, 1, 2},
				5,
			},
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			"2",
			args{
				[]int{10, 1, 2, 7, 6, 1, 5},
				8,
			},
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum2(tt.args.candidates, tt.args.target)
			if !compare(got, tt.want) {
				t.Errorf("combinationSum2() = %v, want %v", got, tt.want)
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
