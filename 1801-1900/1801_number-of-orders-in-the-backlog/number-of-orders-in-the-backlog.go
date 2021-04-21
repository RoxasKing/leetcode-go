package main

import "container/heap"

/*
  You are given a 2D integer array orders, where each orders[i] = [price[i], amount[i], orderType[i]] denotes that amount[i] orders have been placed of type orderType[i] at the price price[i]. The orderType[i] is:
    1. 0 if it is a batch of buy orders, or
    2. 1 if it is a batch of sell orders.
  Note that orders[i] represents a batch of amount[i] independent orders with the same price and order type. All orders represented by orders[i] will be placed before all orders represented by orders[i+1] for all valid i.

  There is a backlog that consists of orders that have not been executed. The backlog is initially empty. When an order is placed, the following happens:
    1. If the order is a buy order, you look at the sell order with the smallest price in the backlog. If that sell order's price is smaller than or equal to the current buy order's price, they will match and be executed, and that sell order will be removed from the backlog. Else, the buy order is added to the backlog.
    2. Vice versa, if the order is a sell order, you look at the buy order with the largest price in the backlog. If that buy order's price is larger than or equal to the current sell order's price, they will match and be executed, and that buy order will be removed from the backlog. Else, the sell order is added to the backlog.
  Return the total amount of orders in the backlog after placing all the orders from the input. Since this number can be large, return it modulo 109 + 7.

  Example 1:
    Input: orders = [[10,5,0],[15,2,1],[25,1,1],[30,4,0]]
    Output: 6
    Explanation: Here is what happens with the orders:
    - 5 orders of type buy with price 10 are placed. There are no sell orders, so the 5 orders are added to the backlog.
    - 2 orders of type sell with price 15 are placed. There are no buy orders with prices larger than or equal to 15, so the 2 orders are added to the backlog.
    - 1 order of type sell with price 25 is placed. There are no buy orders with prices larger than or equal to 25 in the backlog, so this order is added to the backlog.
    - 4 orders of type buy with price 30 are placed. The first 2 orders are matched with the 2 sell orders of the least price, which is 15 and these 2 sell orders are removed from the backlog. The 3rd order is matched with the sell order of the least price, which is 25 and this sell order is removed from the backlog. Then, there are no more sell orders in the backlog, so the 4th order is added to the backlog.
    Finally, the backlog has 5 buy orders with price 10, and 1 buy order with price 30. So the total number of orders in the backlog is 6.

  Example 2:
    Input: orders = [[7,1000000000,1],[15,3,0],[5,999999995,0],[5,1,1]]
    Output: 999999984
    Explanation: Here is what happens with the orders:
    - 109 orders of type sell with price 7 are placed. There are no buy orders, so the 109 orders are added to the backlog.
    - 3 orders of type buy with price 15 are placed. They are matched with the 3 sell orders with the least price which is 7, and those 3 sell orders are removed from the backlog.
    - 999999995 orders of type buy with price 5 are placed. The least price of a sell order is 7, so the 999999995 orders are added to the backlog.
    - 1 order of type sell with price 5 is placed. It is matched with the buy order of the highest price, which is 5, and that buy order is removed from the backlog.
    Finally, the backlog has (1000000000-3) sell orders with price 7, and (999999995-1) buy orders with price 5. So the total number of orders = 1999999991, which is equal to 999999984 % (109 + 7).

  Constraints:
    1. 1 <= orders.length <= 10^5
    2. orders[i].length == 3
    3. 1 <= price[i], amount[i] <= 10^9
    4. orderType[i] is either 0 or 1.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-orders-in-the-backlog
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
