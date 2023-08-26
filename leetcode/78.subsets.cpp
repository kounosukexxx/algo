/*
 * @lc app=leetcode id=78 lang=cpp
 *
 * [78] Subsets
 */
#include <string>
#include <vector>
using namespace std;

// NOTE(encode and decode): subsetを01列で表すことにした。その01列のパターンをrecursionで作り、最後にdecodeした。
// 末尾再帰じゃないから、関数n列分のspaceしようしてる
// currとsubsetで2こvector使ってるからmemoryでbeatできなかった？

// re:
// - cascading: take a elem into consideration one by one
// - backtracking:
//   - 今まで積み上げたans, curr, その他の値をそのまま利用して解いていく(↑自分のも少し入ってる)
//   - 今回は一回最初の数で場合分けして、そのあとbacktracking再帰

// @lc code=start
class Solution {
public:
  void addSubset(vector<vector<int>> &ans, vector<int> nums, vector<bool> curr) {
    vector<int> subset;
    for (int i = 0; i < curr.size(); i++) {
      if (curr[i]) {
        subset.emplace_back(nums[i]);
      }
    }
    ans.emplace_back(subset);
  }

  void recurse(vector<vector<int>> &ans, vector<int> nums, vector<bool> &curr) {
    if (curr.size() == nums.size()) {
      addSubset(ans, nums, curr);
      return;
    }

    curr.emplace_back(false);
    recurse(ans, nums, curr);
    curr.pop_back();

    curr.emplace_back(true);
    recurse(ans, nums, curr);
    curr.pop_back();
  }

  vector<vector<int>> subsets(vector<int> &nums) {
    vector<vector<int>> ans;
    vector<bool> curr;
    recurse(ans, nums, curr);
    return ans;
  }
};
// @lc code=end
