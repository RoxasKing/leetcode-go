package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{4, 2},
			[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.args.n, tt.args.k); !compare(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
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
