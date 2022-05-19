package main

import (
	"reflect"
	"testing"
)

func Test_criticalConnections(t *testing.T) {
	type args struct {
		n           int
		connections [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}}, [][]int{{1, 3}}},
		{"2", args{2, [][]int{{0, 1}}}, [][]int{{0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := criticalConnections(tt.args.n, tt.args.connections); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("criticalConnections() = %v, want %v", got, tt.want)
			}
		})
	}
}
