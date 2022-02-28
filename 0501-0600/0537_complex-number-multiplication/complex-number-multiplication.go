package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

func complexNumberMultiply(num1 string, num2 string) string {
	strs1 := strings.Split(num1, "+")
	r1, _ := strconv.Atoi(strs1[0])
	i1, _ := strconv.Atoi(strs1[1][:len(strs1[1])-1])
	strs2 := strings.Split(num2, "+")
	r2, _ := strconv.Atoi(strs2[0])
	i2, _ := strconv.Atoi(strs2[1][:len(strs2[1])-1])
	return strconv.Itoa(r1*r2-i1*i2) + "+" + strconv.Itoa(r1*i2+r2*i1) + "i"
}
