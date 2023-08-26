/*
 * @lc app=leetcode id=33 lang=cpp
 *
 * [33] Search in Rotated Sorted Array
 */
#include <string>
#include <vector>
using namespace std;

// 方針間違ってる！複雑すぎ！
// @lc code=start
class Solution {
public:
  int recourse(vector<int> &nums, int target, int k, int left, int right) {
    int mid = (left + right) / 2;
    if (nums[mid] == target) {
      return mid;
    }
    if (left >= right) {
      return -1;
    }
    // if (k > nums[left] && target > k) {
    //   return -1;
    // }
    // if (nums[right] < k && target < k) {
    //   return -1;
    // }

    if (target < nums[mid]) {
      // if (nums[left] >= k && target < k) {
      if (nums[left] >= k && target < k) {
        // go right
        return recourse(nums, target, k, mid + 1, nums.size() - 1);
      }
      // go left
      return recourse(nums, target, k, left, mid - 1);
    } else {
      // mid <= target
      // if (nums[right] < k && k < target) {
      //   // go left
      //   return recourse(nums, target, k, 0, mid - 1);
      // }
      if (nums[left] > nums[right]) {
        // k = nums[left];
        // recourse(nums, target, k, left, right);
        if (nums[right] >= target) {
          // go right
          return recourse(nums, target, k, mid + 1, right);
        } else {
          // go left
          return recourse(nums, target, k, left, mid - 1);
        }
      }
      // go right
      return recourse(nums, target, k, mid + 1, right);
    }
  }

  int search(vector<int> &nums, int target) {
    int k = nums[0];
    if (k == target) {
      return 0;
    }
    return recourse(nums, target, k, 0, nums.size() - 1);
  }
};
// @lc code=end
