package main

import "testing"

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{197, 130, 1}}, true},
		{"2", args{[]int{235, 140, 4}}, false},
		{"3", args{[]int{255}}, false},
		{"4", args{[]int{248, 130, 130, 130}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
