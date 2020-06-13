package codinginterviews

/*
  请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

  若队列为空，pop_front 和 max_value 需要返回 -1

  限制：
    1 <= push_back,pop_front,max_value的总操作数 <= 10000
    1 <= value <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotonous Queue

type MaxQueue struct {
	q []int // queue
	m []int // monotonous queue
}

func NewMaxQueue() MaxQueue {
	return MaxQueue{
		q: []int{},
		m: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.m) == 0 {
		return -1
	}
	return this.m[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.m) > 0 && this.m[len(this.m)-1] < value {
		this.m = this.m[:len(this.m)-1]
	}
	this.m = append(this.m, value)
	this.q = append(this.q, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.m) == 0 {
		return -1
	}
	out := this.q[0]
	if this.q[0] == this.m[0] {
		this.m = this.m[1:]
	}
	this.q = this.q[1:]
	return out
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
