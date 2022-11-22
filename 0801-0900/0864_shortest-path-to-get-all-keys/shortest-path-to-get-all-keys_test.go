package main

import "testing"

func Test_shortestPathAllKeys(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"@.a.#", "###.#", "b.A.B"}}, 8},
		{"2", args{[]string{"@..aA", "..B#.", "....b"}}, 6},
		{"3", args{[]string{"@Aa"}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPathAllKeys(tt.args.grid); got != tt.want {
				t.Errorf("shortestPathAllKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
