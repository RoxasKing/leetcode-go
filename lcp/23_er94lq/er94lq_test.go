package main

import "testing"

func Test_isMagic(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{2, 4, 3, 1, 5}}, true},
		{"2", args{[]int{5, 4, 3, 2, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMagic(tt.args.target); got != tt.want {
				t.Errorf("isMagic() = %v, want %v", got, tt.want)
			}
		})
	}
}
