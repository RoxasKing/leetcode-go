package main

// Difficulty:
// Easy

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day, month, year int) string {
	days := day + 365*(year-1971) + (year-1969)/4
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	return week[(days+3)%7]
}
