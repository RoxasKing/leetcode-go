package main

import (
	"math"
)

/*
  There are n cars traveling at different speeds in the same direction along a one-lane road. You are given an array cars of length n, where cars[i] = [positioni, speedi] represents:

  positioni is the distance between the ith car and the beginning of the road in meters. It is guaranteed that positioni < positioni+1.
  speedi is the initial speed of the ith car in meters per second.
  For simplicity, cars can be considered as points moving along the number line. Two cars collide when they occupy the same position. Once a car collides with another car, they unite and form a single car fleet. The cars in the formed fleet will have the same position and the same speed, which is the initial speed of the slowest car in the fleet.

  Return an array answer, where answer[i] is the time, in seconds, at which the ith car collides with the next car, or -1 if the car does not collide with the next car. Answers within 10-5 of the actual answers are accepted.

  Example 1:
    Input: cars = [[1,2],[2,1],[4,3],[7,2]]
    Output: [1.00000,-1.00000,3.00000,-1.00000]
    Explanation: After exactly one second, the first car will collide with the second car, and form a car fleet with speed 1 m/s. After exactly 3 seconds, the third car will collide with the fourth car, and form a car fleet with speed 2 m/s.

  Example 2:
    Input: cars = [[3,4],[5,4],[6,3],[9,1]]
    Output: [2.00000,1.00000,1.50000,-1.00000]

  Constraints:
    1. 1 <= cars.length <= 10^5
    2. 1 <= positioni, speedi <= 10^6
    3. positioni < positioni+1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/car-fleet-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotone Stack
func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)
	out := make([]float64, n)
	for i := range out {
		out[i] = -1.0
	}
	stk := CarStack{&Car{
		cot: math.MaxFloat64,
		pos: float64(cars[n-1][0]),
		spd: cars[n-1][1],
	}}

	for i := n - 2; i >= 0; i-- {
		pos, spd := float64(cars[i][0]), cars[i][1]

		for stk.Len() > 0 && (spd <= stk.Top().spd ||
			(stk.Top().pos-pos)/(float64(spd-stk.Top().spd)) >= stk.Top().cot) {
			stk.Pop()
		}

		cot := math.MaxFloat64
		if len(stk) > 0 {
			cot = (stk.Top().pos - pos) / (float64(spd - stk.Top().spd))
			out[i] = cot
		}

		stk.Push(&Car{
			cot: cot,
			pos: float64(pos),
			spd: spd,
		})
	}

	return out
}

type Car struct {
	cot float64 // collide time
	pos float64 // position
	spd int     // speed
}

type CarStack []*Car

func (cs CarStack) Len() int       { return len(cs) }
func (cs *CarStack) Push(car *Car) { *cs = append(*cs, car) }
func (cs CarStack) Top() *Car      { return cs[cs.Len()-1] }
func (cs *CarStack) Pop() *Car {
	top := cs.Len() - 1
	out := (*cs)[top]
	*cs = (*cs)[:top]
	return out
}
