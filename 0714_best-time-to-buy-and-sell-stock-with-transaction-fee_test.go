package main

import "testing"

func Test_maxProfitVI(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitVI(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitVI() = %v, want %v", got, tt.want)
			}
		})
	}
}
