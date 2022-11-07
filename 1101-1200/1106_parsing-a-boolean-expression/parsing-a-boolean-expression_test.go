package main

import "testing"

func Test_parseBoolExpr(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"!(f)"}, true},
		{"2", args{"|(f,t)"}, false},
		{"3", args{"&(t,f)"}, false},
		{"4", args{"|(&(t,f,t),!(t))"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExpr(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
