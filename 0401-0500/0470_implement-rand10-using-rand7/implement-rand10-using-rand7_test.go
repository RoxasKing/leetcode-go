package main

import (
	"testing"
)

func Test_rand10(t *testing.T) {
	cnt := [11]int{}
	for i := 0; i < 100000; i++ {
		cnt[rand10()]++
	}

	for i := 1; i <= 10; i++ {
		if cnt[i] < 9000 || cnt[i] > 11000 {
			t.Errorf("rand10() test failed!\n")
		}
	}
}
