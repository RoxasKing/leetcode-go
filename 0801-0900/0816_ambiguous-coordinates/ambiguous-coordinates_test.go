package main

import (
	"reflect"
	"testing"
)

func Test_ambiguousCoordinates(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"(123)"}, []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"}},
		{"2", args{"(00011)"}, []string{"(0.001, 1)", "(0, 0.011)"}},
		{"3", args{"(0123)"}, []string{"(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"}},
		{"4", args{"(100)"}, []string{"(10, 0)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ambiguousCoordinates(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ambiguousCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
