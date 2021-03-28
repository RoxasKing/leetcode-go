package main

import "container/heap"

/*
  There is a school that has classes of students and each class will be having a final exam. You are given a 2D integer array classes, where classes[i] = [pass[i], total[i]]. You know beforehand that in the ith class, there are total[i] total students, but only pass[i] number of students will pass the exam.

  You are also given an integer extraStudents. There are another extraStudents brilliant students that are guaranteed to pass the exam of any class they are assigned to. You want to assign each of the extraStudents students to a class in a way that maximizes the average pass ratio across all the classes.

  The pass ratio of a class is equal to the number of students of the class that will pass the exam divided by the total number of students of the class. The average pass ratio is the sum of pass ratios of all the classes divided by the number of the classes.

  Return the maximum possible average pass ratio after assigning the extraStudents students. Answers within 10-5 of the actual answer will be accepted.

  Example 1:
    Input: classes = [[1,2],[3,5],[2,2]], extraStudents = 2
    Output: 0.78333
    Explanation: You can assign the two extra students to the first class. The average pass ratio will be equal to (3/4 + 3/5 + 2/2) / 3 = 0.78333.

  Example 2:
    Input: classes = [[2,4],[3,9],[4,5],[2,10]], extraStudents = 4
    Output: 0.53485

  Constraints:
    1. 1 <= classes.length <= 10^5
    2. classes[i].length == 2
    3. 1 <= pass[i] <= total[i] <= 10^5
    4. 1 <= extraStudents <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-average-pass-ratio
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Priority Queue(Heap Sort)
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	sum := float64(0)
	pq := PriorityQueue{}
	for _, c := range classes {
		if c[0] == c[1] {
			sum += 1.0
		} else {
			heap.Push(&pq, [2]int{c[0], c[1]})
		}
	}
	for extraStudents > 0 && pq.Len() > 0 {
		e := heap.Pop(&pq).([2]int)
		heap.Push(&pq, [2]int{e[0] + 1, e[1] + 1})
		extraStudents--
	}
	for _, class := range pq {
		sum += float64(class[0]) / float64(class[1])
	}
	return sum / float64(len(classes))
}

type PriorityQueue [][2]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	inc1 := float64(pq[i][0]+1)/float64(pq[i][1]+1) - float64(pq[i][0])/float64(pq[i][1])
	inc2 := float64(pq[j][0]+1)/float64(pq[j][1]+1) - float64(pq[j][0])/float64(pq[j][1])
	return inc1 > inc2
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.([2]int)) }
func (pq *PriorityQueue) Pop() interface{} {
	top := len(*pq) - 1
	out := (*pq)[top]
	*pq = (*pq)[:top]
	return out
}
