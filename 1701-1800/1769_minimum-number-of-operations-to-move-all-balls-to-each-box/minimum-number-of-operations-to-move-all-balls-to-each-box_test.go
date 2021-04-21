package main

import (
	"reflect"
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		boxes string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"110"}, []int{1, 1, 3}},
		{"2", args{"001011"}, []int{11, 8, 5, 4, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.boxes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
