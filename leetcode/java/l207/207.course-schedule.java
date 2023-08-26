package l207;

import java.util.*;

// NOTE: 連結リストでやろうとしたけど、int[2]を1単位でやろうとしたら無理やった

/*
 * @lc app=leetcode id=207 lang=java
 *
 * [207] Course Schedule
 */

// @lc code=start
class Solution {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        HashMap<int[], ArrayList<int[]>> list = new HashMap<>(numCourses);
        HashSet<int[]> canTake = new HashSet<>(numCourses);
        HashSet<int[]> visited = new HashSet<>(numCourses);

        for (int[] prerequisite : prerequisites) {
            list.put(prerequisite, new ArrayList<>());
        }

        return false;
    }
}
// @lc code=end
