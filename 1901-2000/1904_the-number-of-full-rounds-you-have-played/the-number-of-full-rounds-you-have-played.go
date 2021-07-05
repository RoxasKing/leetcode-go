package main

import (
	"math"
	"strconv"
	"strings"
)

// Tags:
// Math

func numberOfRounds(startTime string, finishTime string) int {
	s1 := strings.Split(startTime, ":")
	s2 := strings.Split(finishTime, ":")
	sh, _ := strconv.Atoi(s1[0])
	sm, _ := strconv.Atoi(s1[1])
	fh, _ := strconv.Atoi(s2[0])
	fm, _ := strconv.Atoi(s2[1])
	sm = int(math.Ceil(float64(sm)/15.0)) * 15
	if sm == 60 {
		sh++
		sm = 0
	}
	fm = fm / 15 * 15
	if startTime <= finishTime {
		return (fh-sh)*4 + (fm-sm)/15
	}
	return (23-sh)*4 + (60-sm)/15 + fh*4 + fm/15
}
