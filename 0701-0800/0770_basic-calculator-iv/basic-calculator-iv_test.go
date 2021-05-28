package main

import (
	"reflect"
	"testing"
)

func Test_basicCalculatorIV(t *testing.T) {
	type args struct {
		expression string
		evalvars   []string
		evalints   []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"e + 8 - a + 5", []string{"e"}, []int{1}}, []string{"-1*a", "14"}},
		{
			"2",
			args{
				"e - 8 + temperature - pressure",
				[]string{"e", "temperature"},
				[]int{1, 12},
			},
			[]string{"-1*pressure", "5"},
		},
		{"3", args{"(e + 8) * (e - 8)", []string{}, []int{}}, []string{"1*e*e", "-64"}},
		{"4", args{"7 - 7", []string{}, []int{}}, []string{}},
		{"5", args{"a * b * c + b * a * c * 4", []string{}, []int{}}, []string{"5*a*b*c"}},
		{
			"6",
			args{
				"((a - b) * (b - c) + (c - a)) * ((a - b) + (b - c) * (c - a))",
				[]string{},
				[]int{},
			},
			[]string{
				"-1*a*a*b*b",
				"2*a*a*b*c",
				"-1*a*a*c*c",
				"1*a*b*b*b",
				"-1*a*b*b*c",
				"-1*a*b*c*c",
				"1*a*c*c*c",
				"-1*b*b*b*c",
				"2*b*b*c*c",
				"-1*b*c*c*c",
				"2*a*a*b",
				"-2*a*a*c",
				"-2*a*b*b",
				"2*a*c*c",
				"1*b*b*b",
				"-1*b*b*c",
				"1*b*c*c",
				"-1*c*c*c",
				"-1*a*a",
				"1*a*b",
				"1*a*c",
				"-1*b*c",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicCalculatorIV(tt.args.expression, tt.args.evalvars, tt.args.evalints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("basicCalculatorIV() = %v, want %v", got, tt.want)
			}
		})
	}
}
