package main

import "testing"

func Test_trafficCommand(t *testing.T) {
	type args struct {
		directions []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trafficCommand(tt.args.directions); got != tt.want {
				t.Errorf("trafficCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
