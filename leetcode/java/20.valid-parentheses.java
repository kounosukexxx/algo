/*
 * @lc app=leetcode id=20 lang=java
 *
 * [20] Valid Parentheses
 */

import java.util.*;

// @lc code=start
class Solution {
	public boolean isValid(String s) {
		ArrayDeque<Character> stack = new ArrayDeque<>();
		char[] chars = s.toCharArray();
		for (char c : chars) {
			if (c == '(' || c == '{' || c == '[') {
				stack.offerFirst(c);
			} else {
				if (stack.isEmpty()) {
					return false;
				}
				Character popped = stack.pollFirst();
				switch (c) {
					case ')':
						if (popped != '(') {
							return false;
						}
						break;
					case '}':
						if (popped != '{') {
							return false;
						}
						break;
					case ']':
						if (popped != '[') {
							return false;
						}
						break;
					default:
						return false;
				}
			}
		}

		return stack.isEmpty();
	}
}
// @lc code=end
