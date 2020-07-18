package leetcode

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr       []int
		k         int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4}, 3},
		{"2", args{[]int{1, 1, 1, 1, 1}, 1, 0}, 5},
		{"3", args{[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5}, 6},
		{"4", args{[]int{7, 7, 7, 7, 7, 7, 7}, 7, 7}, 1},
		{"5", args{[]int{4, 4, 4, 4}, 4, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr, tt.args.k, tt.args.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
