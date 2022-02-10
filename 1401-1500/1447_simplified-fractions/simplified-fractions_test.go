package main

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{2}, []string{"1/2"}},
		{"2", args{3}, []string{"1/2", "1/3", "2/3"}},
		{"3", args{4}, []string{"1/2", "1/3", "1/4", "2/3", "3/4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifiedFractions(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("simplifiedFractions() = %v, want %v", got, tt.want)
			}
		})
	}
}
