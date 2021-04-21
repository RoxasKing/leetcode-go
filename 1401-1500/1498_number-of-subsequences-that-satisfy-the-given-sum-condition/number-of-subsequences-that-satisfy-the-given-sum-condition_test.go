package main

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{3, 5, 6, 7}, 9}, 4},
		{"2", args{[]int{3, 3, 6, 8}, 10}, 6},
		{"3", args{[]int{2, 3, 3, 4, 6, 7}, 12}, 61},
		{"4", args{[]int{5, 2, 4, 1, 7, 6, 8}, 16}, 127},
		{"5", args{[]int{7, 10, 7, 3, 7, 5, 4}, 12}, 56},
		{
			"6",
			args{
				[]int{14, 4, 6, 6, 20, 8, 5, 6, 8, 12, 6, 10, 14, 9, 17, 16, 9, 7, 14, 11, 14, 15, 13, 11, 10, 18, 13, 17, 17, 14, 17, 7, 9, 5, 10, 13, 8, 5, 18, 20, 7, 5, 5, 15, 19, 14},
				22,
			},
			272187084,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
