/*
 * @lc app=leetcode id=387 lang=java
 *
 * [387] First Unique Character in a String
 */

// @lc code=start

import java.util.HashMap;

class Solution {
	public int firstUniqChar(String s) {
		HashMap<Character, Integer> counter = new HashMap<Character, Integer>();
		for (char c : s.toCharArray()) {
			counter.put(c, counter.getOrDefault(c, 0) + 1);
		}
		for (int i = 0; i < s.length(); i++) {
			if (counter.get(s.charAt(i)) == 1) {
				return i;
			}
		}
		return -1;
	}
}
// @lc code=end
