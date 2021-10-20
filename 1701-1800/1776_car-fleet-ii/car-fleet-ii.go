package main

import (
	"math"
)

// Tags:
// Monotonic Stack
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
