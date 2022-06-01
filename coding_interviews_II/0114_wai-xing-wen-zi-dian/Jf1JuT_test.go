package main

import "testing"

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{[]string{"wrt", "wrf", "er", "ett", "rftt"}}, "wertf"},
		{"2", args{[]string{"z", "x"}}, "zx"},
		{"3", args{[]string{"z", "x", "z"}}, ""},
		{"4", args{[]string{"abc", "ab"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
