package main

import "testing"

func Test_minSpaceWastedKResizing(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{10, 20, 30}, 1}, 10},
		{"2", args{[]int{2, 48, 18, 16, 15, 9, 48, 7, 44, 48}, 1}, 179},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSpaceWastedKResizing(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minSpaceWastedKResizing() = %v, want %v", got, tt.want)
			}
		})
	}
}
