/*
 * @lc app=leetcode id=167 lang=java
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

 // 他mapもできる

import java.util.*;

// @lc code=start
class Solution {
	public int[] twoSum(int[] numbers, int target) {
		for (int i = 0; i < numbers.length; i++) {
			int number = numbers[i];
			if (number * 2 == target) {
				if (numbers[i + 1] == number) {
					return new int[] { i + 1, i + 2 };
				}
			} else {
				int index = Arrays.binarySearch(numbers, target - number);
				if (index >= 0) {
					return new int[] { i + 1, index + 1 };
				}
			}
		}

		return null;
	}
}
// @lc code=end

// this is not effective
// class Solution {
// public int[] twoSum(int[] numbers, int target) {
// List<Integer> list = new ArrayList<>();
// list = Arrays.stream(numbers).boxed().collect(Collectors.toList());
// for (int i = 0; i < numbers.length; i++) {
// int number = numbers[i];
// if (number * 2 == target) {
// if (numbers[i + 1] == number) {
// return new int[] { i + 1, i + 2 };
// }
// } else {
// int index = Collections.binarySearch(list.subList(i + 1, numbers.length),
// target - number);
// if (index >= 0) {
// return new int[] { i + 1, i + 1 + index + 1 };
// }
// }
// }

// return null;
// }
// }
// // @lc code=end