package main

import (
	"reflect"
	"testing"
)

func Test_memLeak(t *testing.T) {
	type args struct {
		memory1 int
		memory2 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{2, 2}, []int{3, 1, 0}},
		{"2", args{8, 11}, []int{6, 0, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := memLeak(tt.args.memory1, tt.args.memory2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memLeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
