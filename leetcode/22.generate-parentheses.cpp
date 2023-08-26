/*
 * @lc app=leetcode id=22 lang=cpp
 *
 * [22] Generate Parentheses
 */
#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  vector<vector<string>> memory;

  vector<string> generateParenthesis(int n) {
    memory = vector<vector<string>>(n);
    for (int i = 1; i <= n; i++) {
      for (int outlineNum = i; outlineNum <= n - 1; outlineNum++) {
        vector<int> inlinePatterns(outlineNum); // TODO:
      }
    }
  };

  vector<string> getParentheses(int num) { return vector<string>(1); }
  void setMemory(vector<int> inlinePatterns) {}
};
// @lc code=end
