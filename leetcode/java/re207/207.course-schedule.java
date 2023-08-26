package re207;

import java.util.*;

/*
 * @lc app=leetcode id=207 lang=java
 *
 * [207] Course Schedule
 */

// @lc code=start
class Solution {
    HashMap<Integer, ArrayList<Integer>> adjList;
    HashSet<Integer> visited = new HashSet<>();
    HashSet<Integer> takableCourses = new HashSet<>();

    public boolean canFinish(int numCourses, int[][] prerequisites) {
        createAdjList(numCourses, prerequisites);

        for (int key : adjList.keySet()) {
            if (!canTake(key)) {
                return false;
            }
        }

        return true;
    }

    void createAdjList(int numCourses, int[][] prerequisites) {
        adjList = new HashMap<>(numCourses);
        for (int[] prerequisite : prerequisites) {
            if (!adjList.containsKey(prerequisite[0])) {
                adjList.put(prerequisite[0], new ArrayList<>());
            }
            adjList.get(prerequisite[0]).add(prerequisite[1]);
        }
    }

    boolean canTake(int key) {
        if (takableCourses.contains(key)) {
            return true;
        }

        ArrayList<Integer> courses = adjList.get(key);
        if (courses == null) {
            takableCourses.add(key);
            return true;
        }

        if (visited.contains(key)) {
            return false;
        }
        visited.add(key);

        for (int i = 0; i < courses.size(); i++) {
            if (!canTake(courses.get(i))) {
                return false;
            }
        }

        visited.remove(key);
        takableCourses.add(key);
        return true;
    }

    // visitedじゃだめ
    // boolean canTake(int key, HashSet<Integer> visited) {
    // if (visited.contains(key)) {
    // return false;
    // }
    // visited.add(key);

    // ArrayList<Integer> courses = adjList.get(key);
    // if (courses == null) {
    // return true;
    // }

    // for (int i = 0; i < courses.size(); i++) {
    // if (!canTake(courses.get(i), visited)) {
    // return false;
    // }
    // }

    // return true;
    // }
}
// @lc code=end
