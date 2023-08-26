/*
 * @lc app=leetcode id=56 lang=java
 *
 * [56] Merge Intervals
 */

import java.util.*;

// @lc code=start
class Solution {
	public int[][] merge(int[][] intervals) {
		if (intervals.length == 1) {
			return intervals;
		}

		Arrays.sort(intervals, new Comparator<int[]>() {
			@Override
			public int compare(int[] a, int[] b) {
				if (a[0] == b[0]) {
					return a[1] - b[1];
				}
				return a[0] - b[0];
			}
		});

		int[] currentInterval = intervals[0];
		List<int[]> ans = new ArrayList<>();
		for (int i = 0; i < intervals.length - 1; i++) {
			if (areOverlapped(currentInterval, intervals[i + 1])) {
				currentInterval = mergeTwo(currentInterval, intervals[i + 1]);
			} else {
				ans.add(currentInterval);
				currentInterval = intervals[i + 1];
			}
			if (i + 1 == intervals.length - 1) {
				ans.add(currentInterval);
			}
		}

		return ans.toArray(new int[ans.size()][]);
	}

	boolean areOverlapped(int[] a, int[] b) {
		return b[0] <= a[1];
	}

	int[] mergeTwo(int[] a, int[] b) {
		return new int[] { a[0], Math.max(a[1], b[1]) };
	}
}
// @lc code=end
