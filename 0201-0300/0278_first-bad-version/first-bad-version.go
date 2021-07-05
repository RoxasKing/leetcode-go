package main

// Tags:
// Binary Search

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

var BadStart int

func isBadVersion(version int) bool {
	return version >= BadStart
}
