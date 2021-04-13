package main

import "testing"

func Test_areSentencesSimilar(t *testing.T) {
	type args struct {
		sentence1 string
		sentence2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"My name is Haley", "My Haley"}, true},
		{"2", args{"of", "A lot of words"}, false},
		{"3", args{"Eating right now", "Eating"}, true},
		{"4", args{"Luky", "Lucccky"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areSentencesSimilar(tt.args.sentence1, tt.args.sentence2); got != tt.want {
				t.Errorf("areSentencesSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
