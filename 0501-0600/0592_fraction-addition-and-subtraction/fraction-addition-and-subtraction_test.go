package main

import "testing"

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"-1/2+1/2"}, "0/1"},
		{"2", args{"-1/2+1/2+1/3"}, "1/3"},
		{"3", args{"1/3-1/2"}, "-1/6"},
		{"4", args{"1/2-4/1-4/3+1/2-5/1"}, "-28/3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionAddition(tt.args.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
