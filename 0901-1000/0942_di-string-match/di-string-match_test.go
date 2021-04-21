package main

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"IDID"}, []int{0, 4, 1, 3, 2}},
		{"2", args{"III"}, []int{0, 1, 2, 3}},
		{"3", args{"DDI"}, []int{3, 2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diStringMatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
