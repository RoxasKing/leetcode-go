package main

import "testing"

func Test_countPairs(t *testing.T) {
	type args struct {
		deliciousness []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 5, 7, 9}}, 4},
		{"2", args{[]int{1, 1, 1, 3, 3, 3, 7}}, 15},
		{"3", args{[]int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.deliciousness); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
