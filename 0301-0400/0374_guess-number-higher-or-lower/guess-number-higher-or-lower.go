package main

// Difficulty:
// Easy

// Tags:
// Binary Search

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)>>1
		res := guess(m)
		if res == 1 {
			l = m + 1
		} else if res == -1 {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

func guess(num int) int {
	if num < PICKED {
		return 1
	} else if num > PICKED {
		return -1
	}
	return 0
}

var PICKED = -1 << 31
