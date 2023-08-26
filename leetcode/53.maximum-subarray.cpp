/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 */
#include <string>
#include <vector>
using namespace std;

// 先に正の数が存在する時のみ考えたのがよかった。
// 一回前処理をして単純にしたのがよかった。

// @lc code=start
class Solution {
  // firstly, slice down left subarray that are added up to negative number.
  // then fix a max total
public:
  int maxSubArray(vector<int> &nums) {
    // when all nums are negative values
    int maxV = -10e4 - 1;
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] >= 0) {
        maxV = 0;
        break;
      }
      maxV = max(maxV, nums[i]);
    }
    if (maxV < 0) {
      return maxV;
    }

    int left = 0;
    int curr = 0;
    int res = 0;
    for (int i = 0; i < nums.size(); ++i) {
      if (curr + nums[i] < 0) {
        if (i != nums.size() - 1) {
          left = i + 1;
        }
        curr = 0;
        continue;
      }
      curr += nums[i];
      res = max(res, curr);
    }

    return res;

    // curr = 0;
    // for (int i = nums.size() - 1; i >= 0; --i) {
    //   if (right <= left) {
    //     break;
    //   }
    //   if (curr + nums[i] < 0) {
    //     if (i != 0) {
    //       right = i - 1;
    //     }
    //     curr = 0;
    //     continue;
    //   }
    //   curr += nums[i];
    // }

    // int res = 0;
    // for (int i = left; i <= right; ++i) {
    //   res += nums[i];
    // }
    // return res;
  };
};
// @lc code=end
