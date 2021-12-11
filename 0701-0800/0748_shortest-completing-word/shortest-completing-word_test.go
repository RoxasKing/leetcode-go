package main

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1s3 PSt", []string{"step", "steps", "stripe", "stepple"}}, "steps"},
		{"2", args{"1s3 456", []string{"looks", "pest", "stew", "show"}}, "pest"},
		{"3", args{"Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}}, "husband"},
		{"4", args{"OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}}, "enough"},
		{"5", args{"iMSlpe4", []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}}, "simple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
