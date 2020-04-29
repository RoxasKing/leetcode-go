package leetcode

import "testing"

func Test_findInMountainArray(t *testing.T) {
	nums1 := []int{3, 5, 3, 2, 0}
	m1 := newMountainArray(len(nums1))
	for i := range nums1 {
		m1.put(nums1[i])
	}
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0, m1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
