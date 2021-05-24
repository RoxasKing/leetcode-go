package main

/*
  You are given a floating-point number hour, representing the amount of time you have to reach the office. To commute to the office, you must take n trains in sequential order. You are also given an integer array dist of length n, where dist[i] describes the distance (in kilometers) of the ith train ride.

  Each train can only depart at an integer hour, so you may need to wait in between each train ride.

    For example, if the 1st train ride takes 1.5 hours, you must wait for an additional 0.5 hours before you can depart on the 2nd train ride at the 2 hour mark.

  Return the minimum positive integer speed (in kilometers per hour) that all the trains must travel at for you to reach the office on time, or -1 if it is impossible to be on time.

  Tests are generated such that the answer will not exceed 107 and hour will have at most two digits after the decimal point.

  Example 1:
    Input: dist = [1,3,2], hour = 6
    Output: 1
    Explanation: At speed 1:
      - The first train ride takes 1/1 = 1 hour.
      - Since we are already at an integer hour, we depart immediately at the 1 hour mark. The second train takes 3/1 = 3 hours.
      - Since we are already at an integer hour, we depart immediately at the 4 hour mark. The third train takes 2/1 = 2 hours.
      - You will arrive at exactly the 6 hour mark.

  Example 2:
    Input: dist = [1,3,2], hour = 2.7
    Output: 3
    Explanation: At speed 3:
      - The first train ride takes 1/3 = 0.33333 hours.
      - Since we are not at an integer hour, we wait until the 1 hour mark to depart. The second train ride takes 3/3 = 1 hour.
      - Since we are already at an integer hour, we depart immediately at the 2 hour mark. The third train takes 2/3 = 0.66667 hours.
      - You will arrive at the 2.66667 hour mark.

  Example 3:
    Input: dist = [1,3,2], hour = 1.9
    Output: -1
    Explanation: It is impossible because the earliest the third train can depart is at the 2 hour mark.

  Constraints:
    1. n == dist.length
    2. 1 <= n <= 10^5
    3. 1 <= dist[i] <= 10^5
    4. 1 <= hour <= 10^9
    5. There will be at most two digits after the decimal point in hour.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-speed-to-arrive-on-time
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

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
