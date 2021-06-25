package main

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}, 6},
		{"2", args{[]string{"8888"}, "0009"}, 1},
		{"3", args{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"}, -1},
		{"4", args{[]string{"0000"}, "8888"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
