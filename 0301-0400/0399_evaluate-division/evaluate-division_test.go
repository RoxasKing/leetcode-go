package main

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			"1",
			args{
				[][]string{{"a", "b"}, {"b", "c"}},
				[]float64{2.0, 3.0},
				[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			"2",
			args{
				[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
				[]float64{1.5, 2.5, 5.0},
				[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			},
			[]float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			"3",
			args{
				[][]string{{"a", "b"}},
				[]float64{0.5},
				[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			},
			[]float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
