package main

// Tags:
// Binary Search

func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	l, r := 1, int(1e7)

	for l < r {
		speed := l + (r-l)>>1
		if calTime(dist, n, speed) > hour {
			l = speed + 1
		} else {
			r = speed
		}
	}

	if calTime(dist, n, l) <= hour {
		return l
	}
	return -1
}

func calTime(dist []int, n, speed int) float64 {
	t := 0
	for _, d := range dist[:n-1] {
		t += d / speed
		if d%speed > 0 {
			t++
		}
	}
	return float64(t) + float64(dist[n-1])/float64(speed)
}
