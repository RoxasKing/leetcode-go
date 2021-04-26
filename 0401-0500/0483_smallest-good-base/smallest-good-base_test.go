package main

import "testing"

func Test_smallestGoodBase(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"13"}, "3"},
		{"2", args{"4681"}, "8"},
		{"3", args{"1000000000000000000"}, "999999999999999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestGoodBase(tt.args.n); got != tt.want {
				t.Errorf("smallestGoodBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
