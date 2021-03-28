package main

import "sort"

/*
  You are given a large sample of integers in the range [0, 255]. Since the sample is so large, it is represented by an array count where count[k] is the number of times that k appears in the sample.

  Calculate the following statistics:
    1. minimum: The minimum element in the sample.
    2. maximum: The maximum element in the sample.
    3. mean: The average of the sample, calculated as the total sum of all elements divided by the total number of elements.
    4. median:
      1. If the sample has an odd number of elements, then the median is the middle element once the sample is sorted.
      2. If the sample has an even number of elements, then the median is the average of the two middle elements once the sample is sorted.
    5. mode: The number that appears the most in the sample. It is guaranteed to be unique.
  Return the statistics of the sample as an array of floating-point numbers [minimum, maximum, mean, median, mode]. Answers within 10-5 of the actual answer will be accepted.

  Example 1:
    Input: count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    Output: [1.00000,3.00000,2.37500,2.50000,3.00000]
    Explanation: The sample represented by count is [1,2,2,2,3,3,3,3].
    The minimum and maximum are 1 and 3 respectively.
    The mean is (1+2+2+2+3+3+3+3) / 8 = 19 / 8 = 2.375.
    Since the size of the sample is even, the median is the average of the two middle elements 2 and 3, which is 2.5.
    The mode is 3 as it appears the most in the sample.

  Example 2:
    Input: count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    Output: [1.00000,4.00000,2.18182,2.00000,1.00000]
    Explanation: The sample represented by count is [1,1,1,1,2,2,2,3,3,4,4].
    The minimum and maximum are 1 and 4 respectively.
    The mean is (1+1+1+1+2+2+2+3+3+4+4) / 11 = 24 / 11 = 2.18181818... (for display purposes, the output shows the rounded number 2.18182).
    Since the size of the sample is odd, the median is the middle element 2.
    The mode is 1 as it appears the most in the sample.

  Constraints:
    1. count.length == 256
    2. 0 <= count[i] <= 10^9
    3. 1 <= sum(count) <= 10^9
    4. The mode of the sample that count represents is unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/statistics-from-a-large-sample
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
