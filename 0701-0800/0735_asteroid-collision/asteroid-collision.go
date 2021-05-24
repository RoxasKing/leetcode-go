package main

/*
  We are given an array asteroids of integers representing asteroids in a row.

  For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

  Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

  Example 1:
    Input: asteroids = [5,10,-5]
    Output: [5,10]
    Explanation: The 10 and -5 collide resulting in 10. The 5 and 10 never collide.

  Example 2:
    Input: asteroids = [8,-8]
    Output: []
    Explanation: The 8 and -8 collide exploding each other.

  Example 3:
    Input: asteroids = [10,2,-5]
    Output: [10]
    Explanation: The 2 and -5 collide resulting in -5. The 10 and -5 collide resulting in 10.

  Example 4:
    Input: asteroids = [-2,-1,1,2]
    Output: [-2,-1,1,2]
    Explanation: The -2 and -1 are moving left, while the 1 and 2 are moving right. Asteroids moving the same direction never meet, so no asteroids will meet each other.

  Constraints:
    1. 2 <= asteroids.length <= 10^4
    2. -1000 <= asteroids[i] <= 1000
    3. asteroids[i] != 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/asteroid-collision
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack

func asteroidCollision(asteroids []int) []int {
	stk := IntStack{}
	for _, a := range asteroids {
		if a > 0 {
			stk.Push(a)
			continue
		}
		ok := true
		for stk.Len() > 0 && stk.Top() > 0 {
			top := stk.Top()
			if top <= -a {
				stk.Pop()
			}
			if top >= -a {
				ok = false
				break
			}
		}
		if ok {
			stk.Push(a)
		}
	}
	return stk
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
