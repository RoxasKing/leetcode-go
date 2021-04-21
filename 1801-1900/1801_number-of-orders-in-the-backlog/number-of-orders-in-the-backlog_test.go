package main

import "testing"

func Test_getNumberOfBacklogOrders(t *testing.T) {
	type args struct {
		orders [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}}, 6},
		{"2", args{[][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}}, 999999984},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfBacklogOrders(tt.args.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
