package main

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{5, 5, 5, 10, 20}}, true},
		{"2", args{[]int{5, 5, 10}}, true},
		{"3", args{[]int{10, 10}}, false},
		{"4", args{[]int{5, 5, 10, 10, 20}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
