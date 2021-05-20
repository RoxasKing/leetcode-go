package main

import (
	"reflect"
	"testing"
)

func Test_numMovesStones(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{1, 2, 5}, []int{1, 2}},
		{"2", args{4, 3, 2}, []int{0, 0}},
		{"3", args{3, 5, 1}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMovesStones(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
