package main

// Tags:
// Binary Search

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	out := [][]int{}
	for i := 1; i <= 1000; i++ {
		l, r := 1, 1000
		for l <= r {
			j := (l + r) >> 1
			res := customFunction(i, j)
			if res < z {
				l = j + 1
			} else if res > z {
				r = j - 1
			} else {
				out = append(out, []int{i, j})
				break
			}
		}
	}
	return out
}
