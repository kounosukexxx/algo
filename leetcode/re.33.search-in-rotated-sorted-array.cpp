/*
 * @lc app=leetcode id=33 lang=cpp
 *
 * [33] Search in Rotated Sorted Array
 */
#include <string>
#include <vector>
using namespace std;

// NOTE: O(2*X)なら問題ないケース。binary searchを2回する。
// @lc code=start
class Solution {
public:
  int getMaxIndex(vector<int> nums, int left, int right) {
    if (nums[left] <= nums[right]) {
      // right is max
      return right;
    }

    int mid = (left + right) / 2;
    if (nums[mid] > nums[mid + 1]) {
      return mid;
    }

    if (nums[mid] > nums[right]) {
      return getMaxIndex(nums, mid + 1, right);
    } else {
      return getMaxIndex(nums, left, mid - 1);
    }
  }
  int binarySearch(vector<int> nums, int target, int left, int right) {
    if (left <= right) {
      int mid = (left + right) / 2;
      if (nums[mid] == target) {
        return mid;
      }
      if (left == right) {
        return -1;
      }

      if (target < nums[mid]) {
        return binarySearch(nums, target, left, mid - 1);
      } else {
        return binarySearch(nums, target, mid + 1, right);
      }
    }

    return -1;
  }

  int search(vector<int> &nums, int target) {
    int k = nums[0];
    int maxIndex = getMaxIndex(nums, 0, nums.size() - 1);
    if (target >= k) {
      return binarySearch(nums, target, 0, maxIndex);
    } else {
      return binarySearch(nums, target, maxIndex + 1, nums.size() - 1);
    }
  };
};
// @lc code=end
