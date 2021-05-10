package main

import "testing"

func Test_maximumPopulation(t *testing.T) {
	type args struct {
		logs [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[][]int{{1993, 1999}, {2000, 2010}}}, 1993},
		{"2", args{[][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}}, 1960},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPopulation(tt.args.logs); got != tt.want {
				t.Errorf("maximumPopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}
