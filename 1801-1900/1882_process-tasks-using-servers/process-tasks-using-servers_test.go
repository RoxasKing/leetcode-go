package main

import (
	"reflect"
	"testing"
)

func Test_assignTasks(t *testing.T) {
	type args struct {
		servers []int
		tasks   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{3, 3, 2}, []int{1, 2, 3, 2, 1, 2}}, []int{2, 2, 0, 2, 1, 2}},
		{"2", args{[]int{5, 1, 4, 3, 2}, []int{2, 1, 2, 4, 5, 2, 1}}, []int{1, 4, 1, 4, 1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignTasks(tt.args.servers, tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("assignTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
