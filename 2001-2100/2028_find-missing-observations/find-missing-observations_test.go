package main

import (
	"reflect"
	"testing"
)

func Test_missingRolls(t *testing.T) {
	type args struct {
		rolls []int
		mean  int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 4, 3}, 4, 2}, []int{6, 6}},
		{"2", args{[]int{1, 5, 6}, 3, 4}, []int{3, 2, 2, 2}},
		{"3", args{[]int{1, 2, 3, 4}, 6, 4}, []int{}},
		{"4", args{[]int{6, 3, 4, 3, 5, 3}, 1, 6}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingRolls(tt.args.rolls, tt.args.mean, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("missingRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
