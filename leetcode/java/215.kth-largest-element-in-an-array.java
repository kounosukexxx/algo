/*
 * @lc app=leetcode id=215 lang=java
 *
 * [215] Kth Largest Element in an Array
 */

import java.util.*;

// @lc code=start
class Solution {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> maxHeap = new PriorityQueue<>(nums.length, Collections.reverseOrder());
        for (Integer num : nums) {
            maxHeap.offer(num);
        }

        Integer polled = null;
        for (int i = 0; i < k; i++) {
            polled = maxHeap.poll();
        }

        return polled;
    }
}
// @lc code=end
