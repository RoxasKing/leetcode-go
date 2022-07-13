package main

import "testing"

func Test_evaluate(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"(let x 2 (mult x (let x 3 y 4 (add x y))))"}, 14},
		{"2", args{"(let x 3 x 2 x)"}, 2},
		{"3", args{"(let x 1 y 2 x (add x y) (add x y))"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.expression); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
