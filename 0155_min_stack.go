package My_LeetCode_In_Go

type MinStack struct {
	stackArray []int
	helpArray  []int
}

/** initialize your data structure here. */
func NewMinStack() MinStack {
	return MinStack{stackArray: make([]int, 0), helpArray: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.stackArray = append(this.stackArray, x)
	if len(this.helpArray) == 0 ||
		x <= this.helpArray[len(this.helpArray)-1] {
		this.helpArray = append(this.helpArray, x)
	}
}

func (this *MinStack) Pop() {
	if this.stackArray[len(this.stackArray)-1] ==
		this.helpArray[len(this.helpArray)-1] {
		this.helpArray = this.helpArray[:len(this.helpArray)-1]
	}
	this.stackArray = this.stackArray[:len(this.stackArray)-1]

}

func (this *MinStack) Top() int {
	return this.stackArray[len(this.stackArray)-1]
}

func (this *MinStack) GetMin() int {
	return this.helpArray[len(this.helpArray)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
