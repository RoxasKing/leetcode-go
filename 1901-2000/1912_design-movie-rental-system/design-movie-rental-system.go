package main

import (
	"container/heap"
	"sort"
)

// Important!

// Priority Queue

type MovieRentingSystem struct {
	movInf  map[int]*movieInfo
	movieQ  map[int]*movieQueue
	rented  map[int]bool
	rentedQ *rentedQueue
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	movInf := map[int]*movieInfo{}
	movieQ := map[int]*movieQueue{}
	for _, e := range entries {
		s, m, p := e[0], e[1], e[2]
		info := &movieInfo{shop: s, movie: m, price: p}
		movInf[s*MOD+m] = info
		if _, ok := movieQ[m]; !ok {
			movieQ[m] = &movieQueue{}
		}
		heap.Push(movieQ[m], info)
	}
	return MovieRentingSystem{
		movInf:  movInf,
		movieQ:  movieQ,
		rented:  map[int]bool{},
		rentedQ: &rentedQueue{},
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	if _, ok := this.movieQ[movie]; !ok {
		return nil
	}
	arr := []*movieInfo{}
	for len(arr) < 5 && this.movieQ[movie].Len() > 0 {
		info := heap.Pop(this.movieQ[movie]).(*movieInfo)
		key := info.shop*MOD + info.movie
		if this.rented[key] {
			continue
		}
		if len(arr) > 0 && *arr[len(arr)-1] == *info {
			continue
		}
		arr = append(arr, info)
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].price != arr[j].price {
			return arr[i].price < arr[j].price
		}
		return arr[i].shop < arr[j].shop
	})
	out := make([]int, 0, 5)
	for _, e := range arr {
		out = append(out, e.shop)
		heap.Push(this.movieQ[movie], e)
	}
	return out
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	key := shop*MOD + movie
	this.rented[key] = true
	movieInfo := this.movInf[key]
	heap.Push(this.rentedQ, movieInfo)
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	key := shop*MOD + movie
	this.rented[key] = false
	movieInfo := this.movInf[key]
	heap.Push(this.movieQ[movie], movieInfo)
}

func (this *MovieRentingSystem) Report() [][]int {
	arr := []*movieInfo{}
	for len(arr) < 5 && this.rentedQ.Len() > 0 {
		info := heap.Pop(this.rentedQ).(*movieInfo)
		key := info.shop*MOD + info.movie
		if !this.rented[key] {
			continue
		}
		if len(arr) > 0 && *arr[len(arr)-1] == *info {
			continue
		}
		arr = append(arr, info)
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].price != arr[j].price {
			return arr[i].price < arr[j].price
		} else if arr[i].shop != arr[j].shop {
			return arr[i].shop < arr[j].shop
		}
		return arr[i].movie < arr[j].movie
	})
	out := make([][]int, 0, 5)
	for _, e := range arr {
		out = append(out, []int{e.shop, e.movie})
		heap.Push(this.rentedQ, e)
	}
	return out
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */

var MOD = 100001

type movieInfo struct {
	shop, movie, price int
}

type movieQueue []*movieInfo

func (q movieQueue) Len() int { return len(q) }
func (q movieQueue) Less(i, j int) bool {
	if q[i].price != q[j].price {
		return q[i].price < q[j].price
	}
	return q[i].shop < q[j].shop
}
func (q movieQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *movieQueue) Push(x interface{}) { *q = append(*q, x.(*movieInfo)) }
func (q *movieQueue) Pop() interface{} {
	top := q.Len() - 1
	out := (*q)[top]
	*q = (*q)[:top]
	return out
}

type rentedQueue []*movieInfo

func (q rentedQueue) Len() int { return len(q) }
func (q rentedQueue) Less(i, j int) bool {
	if q[i].price != q[j].price {
		return q[i].price < q[j].price
	} else if q[i].shop != q[j].shop {
		return q[i].shop < q[j].shop
	}
	return q[i].movie < q[j].movie
}
func (q rentedQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *rentedQueue) Push(x interface{}) { *q = append(*q, x.(*movieInfo)) }
func (q *rentedQueue) Pop() interface{} {
	top := q.Len() - 1
	out := (*q)[top]
	*q = (*q)[:top]
	return out
}
