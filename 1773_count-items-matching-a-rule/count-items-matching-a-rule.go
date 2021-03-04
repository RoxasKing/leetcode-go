package main

/*
  You are given an array items, where each items[i] = [typei, colori, namei] describes the type, color, and name of the ith item. You are also given a rule represented by two strings, ruleKey and ruleValue.

  The ith item is said to match the rule if one of the following is true:
    1. ruleKey == "type" and ruleValue == typei.
    2. ruleKey == "color" and ruleValue == colori.
    3. ruleKey == "name" and ruleValue == namei.
  Return the number of items that match the given rule.

  Example 1:
    Input: items = [["phone","blue","pixel"],["computer","silver","lenovo"],["phone","gold","iphone"]], ruleKey = "color", ruleValue = "silver"
    Output: 1
    Explanation: There is only one item matching the given rule, which is ["computer","silver","lenovo"].

  Example 2:
    Input: items = [["phone","blue","pixel"],["computer","silver","phone"],["phone","gold","iphone"]], ruleKey = "type", ruleValue = "phone"
    Output: 2
    Explanation: There are only two items matching the given rule, which are ["phone","blue","pixel"] and ["phone","gold","iphone"]. Note that the item ["computer","silver","phone"] does not match.

  Constraints:
    1. 1 <= items.length <= 10^4
    2. 1 <= typei.length, colori.length, namei.length, ruleValue.length <= 10
    3. ruleKey is equal to either "type", "color", or "name".
    4. All strings consist only of lowercase letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-items-matching-a-rule
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	out := 0
	for _, item := range items {
		if ruleKey == "type" && item[0] == ruleValue ||
			ruleKey == "color" && item[1] == ruleValue ||
			ruleKey == "name" && item[2] == ruleValue {
			out++
		}
	}
	return out
}
