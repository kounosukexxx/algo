/*
 * @lc app=leetcode id=937 lang=java
 *
 * [937] Reorder Data in Log Files
 */

import java.util.*;

// @lc code=start
class Solution {
	public String[] reorderLogFiles(String[] logs) {
		ArrayList<String> ans = new ArrayList<String>();
		ArrayList<String> digitLogs = new ArrayList<String>();
		for (String log : logs) {
			if (isDigitLog(log)) {
				digitLogs.add(log);
			} else {
				ans.add(log);
			}
		}

		// sort
		Collections.sort(ans, new Comparator<String>() {
			@Override
			public int compare(String s1, String s2) {
				String[] s1Split = s1.split(" ", 2);
				String[] s2Split = s2.split(" ", 2);
				int cmp = s1Split[1].compareTo(s2Split[1]);
				if (cmp != 0) {
					return cmp;
				}
				return s1Split[0].compareTo(s2Split[0]);
			}
		});

		for (String log : digitLogs) {
			ans.add(log);
		}

		return ans.toArray(new String[ans.size()]);
	}

	boolean isDigitLog(String log) {
		String[] splitted = log.split(" ", 2);
		return splitted[1].charAt(0) >= '0' && splitted[1].charAt(0) <= '9';
	}
}
// @lc code=end
