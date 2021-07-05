package main

// Tags:
// Two Pointers
func isHappy(n int) bool {
	slow := getNext(n)
	fast := getNext(slow)
	for fast != 1 {
		if slow == fast {
			return false
		}
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return true
}

func getNext(n int) int {
	var res int
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}

func isHappy2(n int) bool {
	dict := map[int]struct{}{
		4:   {},
		16:  {},
		37:  {},
		58:  {},
		89:  {},
		145: {},
		42:  {},
		20:  {},
	}
	for {
		if n == 1 {
			return true
		}
		if _, ok := dict[n]; ok {
			return false
		}
		var newN int
		for n != 0 {
			newN += (n % 10) * (n % 10)
			n /= 10
		}
		n = newN
	}
}
