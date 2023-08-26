/*
 * @lc app=leetcode id=216 lang=cpp
 *
 * [216] Combination Sum III
 */
#include <string>
#include <vector>
using namespace std;
// recurse:
// ↓の再帰の使い方がよい。再帰つくるなら、まずここ作って見通しつけるのいいかも
//   逆にこれができないと無理
// curr.emplace_back(i);
// recurse(ans, curr, k - 1, n - i);
// curr.pop_back();

// @lc code=start
class Solution {
public:
  void recurse(vector<vector<int>> &ans, vector<int> &curr, int k, int n) {
    for (int i = 1; i <= 9; i++) {
      if (k == 0) {
        if (n == 0) {
          ans.emplace_back(curr);
        }
        return;
      }
      if (curr.size() > 0 && i <= curr[curr.size() - 1]) {
        continue;
      }
      curr.emplace_back(i);
      recurse(ans, curr, k - 1, n - i);
      curr.pop_back();
    }
    return;
  }

  vector<vector<int>> combinationSum3(int k, int n) {
    vector<vector<int>> ans;
    vector<int> curr;
    recurse(ans, curr, k, n);
    return ans;
  };
};
// @lc code=end
