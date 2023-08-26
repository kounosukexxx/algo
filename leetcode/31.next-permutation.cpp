/*
 * @lc app=leetcode id=31 lang=cpp
 *
 * [31] Next Permutation
 */
#include <string>
#include <vector>
using namespace std;

// NOTE: 方針がむずい
// sort:
// 辞書式: 一番でかいのが、降順、一番小さいのは昇順。
//   右から左に見ていって、昇順でなくなった点が、今回大事な”小改善”のポイント。
//   それで完成と思いきや、その時の値は欲しい値でなく、そこから右をソートしたものが答え
//   - ⭐️「変更を加えたら、今までの事象は崩れる」
//   - 63525321
// @lc code=start
class Solution {
  // void swap(vector<int> &nums, int i, int j) {
  //   int tmp = nums[i];
  //   nums[i] = nums[j];
  //   nums[j] = tmp;
  // }

  void sortAscending(vector<int> &nums, int left) {
    sort(nums.begin() + left, nums.end());
  }

public:
  void nextPermutation(vector<int> &nums) {
    for (int i = nums.size() - 1; i > -1; i--) {
      for (int j = nums.size() - 1; j > i; j--) {
        if (nums[j] > nums[i]) {
          swap(nums[i], nums[j]);
          if (i != nums.size() - 1) {
            sortAscending(nums, i + 1);
          }
          return;
        }
      }
    }

    for (int i = 0; i < nums.size() / 2; i++) {
      swap(nums[i], nums[nums.size() - 1 - i]);
    }
  }
};
// @lc code=end
