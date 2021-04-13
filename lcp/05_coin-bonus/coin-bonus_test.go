package main

import (
	"reflect"
	"testing"
)

func Test_bonus(t *testing.T) {
	type args struct {
		n          int
		leadership [][]int
		operations [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				6,
				[][]int{{1, 2}, {1, 6}, {2, 3}, {2, 5}, {1, 4}},
				[][]int{{1, 1, 500}, {2, 2, 50}, {3, 1}, {2, 6, 15}, {3, 1}},
			},
			[]int{650, 665},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bonus(tt.args.n, tt.args.leadership, tt.args.operations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bonus() = %v, want %v", got, tt.want)
			}
		})
	}
}
