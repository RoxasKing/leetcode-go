package main

import "container/heap"

// Priority Queue(Heap Sort)
func getNumberOfBacklogOrders(orders [][]int) int {
	sellQ := MinHeap{}
	buyQ := MaxHeap{}
	for _, order := range orders {
		price, amount, typ := order[0], order[1], order[2]
		if typ == 0 {
			for sellQ.Len() > 0 && sellQ[0][0] <= price && amount > 0 {
				if sellQ[0][1] <= amount {
					amount -= sellQ[0][1]
					sellQ[0][1] = 0
					heap.Pop(&sellQ)
				} else {
					sellQ[0][1] -= amount
					amount = 0
				}
			}
			if amount > 0 {
				heap.Push(&buyQ, [2]int{price, amount})
			}
		} else {
			for buyQ.Len() > 0 && buyQ[0][0] >= price && amount > 0 {
				if buyQ[0][1] <= amount {
					amount -= buyQ[0][1]
					buyQ[0][1] = 0
					heap.Pop(&buyQ)
				} else {
					buyQ[0][1] -= amount
					amount = 0
				}
			}
			if amount > 0 {
				heap.Push(&sellQ, [2]int{price, amount})
			}
		}
	}

	out := 0
	for _, order := range sellQ {
		out = (out + order[1]) % (1e9 + 7)
	}
	for _, order := range buyQ {
		out = (out + order[1]) % (1e9 + 7)
	}
	return out
}

type MinHeap [][2]int

func (mh MinHeap) Len() int            { return len(mh) }
func (mh MinHeap) Less(i, j int) bool  { return mh[i][0] < mh[j][0] }
func (mh MinHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MinHeap) Push(x interface{}) { *mh = append(*mh, x.([2]int)) }
func (mh *MinHeap) Pop() interface{} {
	top := mh.Len() - 1
	out := (*mh)[top]
	*mh = (*mh)[:top]
	return out
}

type MaxHeap [][2]int

func (mh MaxHeap) Len() int            { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool  { return mh[i][0] > mh[j][0] }
func (mh MaxHeap) Swap(i, j int)       { mh[i], mh[j] = mh[j], mh[i] }
func (mh *MaxHeap) Push(x interface{}) { *mh = append(*mh, x.([2]int)) }
func (mh *MaxHeap) Pop() interface{} {
	top := mh.Len() - 1
	out := (*mh)[top]
	*mh = (*mh)[:top]
	return out
}
