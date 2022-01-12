package main

import "testing"

func Test_slowestKey(t *testing.T) {
	type args struct {
		releaseTimes []int
		keysPressed  string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"1", args{[]int{9, 29, 49, 50}, "cbcd"}, 'c'},
		{"2", args{[]int{12, 23, 36, 46, 62}, "spuda"}, 'a'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slowestKey(tt.args.releaseTimes, tt.args.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
