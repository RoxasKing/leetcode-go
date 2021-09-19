package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"123", 6}, []string{"1*2*3", "1+2+3"}},
		{"2", args{"232", 8}, []string{"2*3+2", "2+3*2"}},
		{"3", args{"105", 5}, []string{"1*0+5", "10-5"}},
		{"4", args{"00", 0}, []string{"0*0", "0+0", "0-0"}},
		{"5", args{"3456237490", 9191}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addOperators(tt.args.num, tt.args.target)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
