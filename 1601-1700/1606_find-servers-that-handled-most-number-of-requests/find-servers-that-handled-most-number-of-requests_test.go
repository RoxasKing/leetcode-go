package main

import (
	"reflect"
	"testing"
)

func Test_busiestServers(t *testing.T) {
	type args struct {
		k       int
		arrival []int
		load    []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{"1", args{3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}}, []int{1}},
		{"2", args{3, []int{1, 2, 3, 4}, []int{1, 2, 1, 2}}, []int{0}},
		{"3", args{3, []int{1, 2, 3}, []int{10, 12, 11}}, []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := busiestServers(tt.args.k, tt.args.arrival, tt.args.load); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("busiestServers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
