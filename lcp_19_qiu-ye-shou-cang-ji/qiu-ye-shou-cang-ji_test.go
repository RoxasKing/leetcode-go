package main

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		leaves string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"rrryyyrryyyrr"}, 2},
		{"2", args{"ryr"}, 0},
		{
			"3",
			args{"ryyryyyrryyyyyryyyrrryyyryryyyyryyrrryryyyryrryrrrryyyrrrrryryyrrrrryyyryyryrryryyryyyyryyrryrryryy"},
			41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperations(tt.args.leaves); got != tt.want {
				t.Errorf("minimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
