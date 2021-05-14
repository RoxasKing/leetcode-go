package main

import "testing"

func Test_findString(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta"}, -1},
		{"2", args{[]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ball"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findString(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("findString() = %v, want %v", got, tt.want)
			}
		})
	}
}
