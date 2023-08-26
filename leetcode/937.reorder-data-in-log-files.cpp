// 途中でjavaの方が便利ってなった

/*
 * @lc app=leetcode id=937 lang=cpp
 *
 * [937] Reorder Data in Log Files
 */
#include <sort>
#include <string>
#include <vector>
using namespace std;

// @lc code=start
class Solution {
public:
  vector<string> reorderLogFiles(vector<string> &logs) {
    vector<string> ans, digitLogs;
    for (auto &log : logs) {
      if (isDigitLog(log)) {
        digitLogs.emplace_back(log);
      } else {
        ans.emplace_back(log);
      }
    }

    // sort(ans.begin(), ans.end(), )
  }

  int getContentBeginIndex(string &log) { return log.find(' '); }

  bool isDigitLog(string &log) { return false; }

  bool asc(string &a, string &b) {}
};
// @lc code=end
