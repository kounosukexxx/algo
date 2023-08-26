/*
 * @lc app=leetcode id=39 lang=cpp
 *
 * [39] Combination Sum
 */

#include <string>
#include <unordered_set>
#include <vector>
using namespace std;

// 再帰の時、値渡しの変数はどこからきているか注意すること。forのスコープが消滅したら、
// スコープ前の時の値に戻ることに注意

// recursion is difficult: 再帰で頑張ろうとしたけど、配列を追加し続けるのが地獄で無理だったパターン。
// 頭では説明できるが、配列とかは「プログラミング固有」のことなので、書いてみるときつかった。
// - と思ったが、再帰の使い方がかなり悪かった（見通し悪いやり方だった）
// @lc code=start
class Solution {
public:
  void recourse(vector<vector<int>> &ans, unordered_set<int> &set,
                vector<int> &curr, vector<int> &candidates, int target) {
    for (int i = 0; i < candidates.size(); i++) {
      if (curr.size() != 0 && candidates[i] < curr[curr.size() - 1]) {
        // 小さいのを追加しようとしたとき。
        continue;
      }

      if (candidates[i] == target) {
        curr.push_back(candidates[i]);
        ans.push_back(curr);
        curr.pop_back();
        if (curr.size() > 0) {
          curr.pop_back(); // those case that the ans length is not 1
        }
        return;
      }

      if (candidates[i] > target) {
        curr.pop_back();
        return;
      }

      // initial value以外で、同じtargetは詮索しない
      // if (set.find(target) != set.end()) {
      //   curr.pop_back();
      //   return;
      // } else {
      //   curr.push_back(candidates[i]);
      //   set.insert(target);
      // }

      recourse(ans, set, curr, candidates, target - candidates[i]);
    }
  }

  vector<vector<int>> combinationSum(vector<int> &candidates, int target) {
    vector<vector<int>> ans;
    vector<int> curr;
    unordered_set<int> set;
    recourse(ans, set, curr, candidates, target);
    return ans;
  }
};
// @lc code=end
