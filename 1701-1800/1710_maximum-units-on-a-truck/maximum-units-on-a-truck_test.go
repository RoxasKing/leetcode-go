package main

import "testing"

func Test_maximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4}, 8},
		{"2", args{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10}, 91},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("maximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
