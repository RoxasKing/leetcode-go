package codinginterviews

// ListNode : Linked list data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode : Binary tree data structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Max : Returns the largest number in a and b
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min : Returns the smallest  number in a and b
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs : Returns the absolute value of n
func Abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}

// Gcd : Returns greatest common divisor of a and b
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Pow returns a**b, the base-a exponential of b.
func Pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

type Stack struct {
	elems []int
}

func NewStack() *Stack {
	return &Stack{
		elems: []int{},
	}
}

func (s *Stack) Size() int {
	return len(s.elems)
}

func (s *Stack) Push(value int) {
	s.elems = append(s.elems, value)
}

func (s *Stack) Peek() int {
	if s.Size() == 0 {
		return -1
	}
	return s.elems[len(s.elems)-1]
}

func (s *Stack) Pop() int {
	if s.Size() == 0 {
		return -1
	}
	value := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]
	return value
}
