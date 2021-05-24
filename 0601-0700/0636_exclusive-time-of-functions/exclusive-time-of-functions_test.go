package main

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	type args struct {
		n    int
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}}, []int{3, 4}},
		{"2", args{1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}}, []int{8}},
		{"3", args{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}}, []int{7, 1}},
		{"4", args{2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"}}, []int{8, 1}},
		{"5", args{1, []string{"0:start:0", "0:end:0"}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.args.n, tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
