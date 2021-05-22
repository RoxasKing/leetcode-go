package main

/*
  We are given non-negative integers nums[i] which are written on a chalkboard.  Alice and Bob take turns erasing exactly one number from the chalkboard, with Alice starting first.  If erasing a number causes the bitwise XOR of all the elements of the chalkboard to become 0, then that player loses.  (Also, we'll say the bitwise XOR of one element is that element itself, and the bitwise XOR of no elements is 0.)

  Also, if any player starts their turn with the bitwise XOR of all the elements of the chalkboard equal to 0, then that player wins.

  Return True if and only if Alice wins the game, assuming both players play optimally.

  Example:
    Input: nums = [1, 1, 2]
    Output: false
    Explanation:
      Alice has two choices: erase 1 or erase 2.
      If she erases 1, the nums array becomes [1, 2]. The bitwise XOR of all the elements of the chalkboard is 1 XOR 2 = 3. Now Bob can remove any element he wants, because Alice will be the one to erase the last element and she will lose.
      If Alice erases 2 first, now nums becomes [1, 1]. The bitwise XOR of all the elements of the chalkboard is 1 XOR 1 = 0. Alice will lose.

  Notes:
    1. 1 <= N <= 1000.
    2. 0 <= nums[i] <= 2^16.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/chalkboard-xor-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math
// Bit Manipulation

func xorGame(nums []int) bool {
	if len(nums)&1 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}
