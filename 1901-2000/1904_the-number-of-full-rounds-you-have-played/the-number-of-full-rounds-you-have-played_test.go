package main

import "testing"

func Test_numberOfRounds(t *testing.T) {
	type args struct {
		startTime  string
		finishTime string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"12:01", "12:44"}, 1},
		{"2", args{"20:00", "06:00"}, 40},
		{"3", args{"00:00", "23:59"}, 95},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfRounds(tt.args.startTime, tt.args.finishTime); got != tt.want {
				t.Errorf("numberOfRounds() = %v, want %v", got, tt.want)
			}
		})
	}
}
