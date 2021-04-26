package main

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{123},
			"One Hundred Twenty Three",
		},
		{
			"2",
			args{12345},
			"Twelve Thousand Three Hundred Forty Five",
		},
		{
			"3",
			args{1234567},
			"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			"4",
			args{1234567891},
			"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
		{
			"5",
			args{0},
			"Zero",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
