package main

func lemonadeChange(bills []int) bool {
	count5, count10 := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			count5++
		} else if bill == 10 {
			if count5 == 0 {
				return false
			}
			count5--
			count10++
		} else {
			if count10 > 0 {
				if count5 == 0 {
					return false
				}
				count5--
				count10--
			} else {
				if count5 < 3 {
					return false
				}
				count5 -= 3
			}
		}
	}
	return true
}
