package main

import (
	"strconv"
	"strings"
)

// Dfficulty:
// Easy

func dayOfYear(date string) int {
	isLeapYear := func(year int) bool { return year%4 == 0 && year%100 != 0 || year%400 == 0 }
	days := [2][12]int{
		{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
		{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	}
	strs := strings.Split(date, "-")
	year, _ := strconv.Atoi(strs[0])
	month, _ := strconv.Atoi(strs[1])
	day, _ := strconv.Atoi(strs[2])
	idx := 0
	if isLeapYear(year) {
		idx++
	}
	out := day
	for i := 0; i < month-1; i++ {
		out += days[idx][i]
	}
	return out
}
