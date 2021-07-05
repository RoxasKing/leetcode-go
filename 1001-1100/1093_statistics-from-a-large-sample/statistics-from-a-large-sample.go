package main

import "sort"

func sampleStats(count []int) []float64 {
	minimum := -1
	maximum := -1
	mode := 0
	sum := 0
	cnt := 0
	arr := []int{}
	for i := range count {
		if count[i] == 0 {
			continue
		}
		if minimum == -1 {
			minimum = i
		}
		maximum = i
		if count[i] > count[mode] {
			mode = i
		}
		sum += count[i] * i
		cnt += count[i]
		arr = append(arr, i)
	}
	mean := float64(sum) / float64(cnt)
	median := float64(0)
	sort.Ints(arr)
	cntL := 0
	a, b := 0, 0
	for i, num := range arr {
		cntL += count[num]
		if cnt&1 == 1 {
			if cntL > cnt>>1 {
				a, b = num, num
				break
			}
		} else {
			if cntL == cnt>>1 {
				a, b = num, arr[i+1]
				break
			} else if cntL > cnt>>1 {
				a, b = num, num
				break
			}
		}
	}
	median = float64(a+b) / 2.0
	return []float64{float64(minimum), float64(maximum), mean, median, float64(mode)}
}
