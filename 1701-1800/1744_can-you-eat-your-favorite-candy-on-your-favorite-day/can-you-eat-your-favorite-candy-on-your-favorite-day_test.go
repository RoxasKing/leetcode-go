package main

import (
	"reflect"
	"testing"
)

func Test_canEat(t *testing.T) {
	type args struct {
		candiesCount []int
		queries      [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			"1",
			args{
				[]int{7, 4, 5, 3, 8},
				[][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 1000000000}},
			},
			[]bool{true, false, true},
		},
		{
			"2",
			args{
				[]int{5, 2, 6, 4, 1},
				[][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}},
			},
			[]bool{false, true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.args.candiesCount, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
