package main

import "testing"

func Test_partitionDisjoint(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{5, 0, 3, 8, 6}}, 3},
		{"2", args{[]int{1, 1, 1, 0, 6, 12}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partitionDisjoint(tt.args.nums); got != tt.want {
				t.Errorf("partitionDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
