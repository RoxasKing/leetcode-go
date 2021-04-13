package main

import (
	"reflect"
	"testing"
)

func Test_fraction(t *testing.T) {
	type args struct {
		cont []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 2, 0, 2}}, []int{13, 4}},
		{"2", args{[]int{0, 0, 3}}, []int{3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fraction(tt.args.cont); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
