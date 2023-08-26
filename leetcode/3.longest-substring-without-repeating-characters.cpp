/*
 * @lc app=leetcode id=3 lang=cp
 *
 * [3] Longest Substring Without Repeating Characters
 */

#include <string>
#include <unordered_map>
using namespace std;

// NOTE: maxLen = max(maxLen, len);をやるのは計算が終わった後
// @lc code=start
class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    int maxLen = 0;
    int len;
    for (int i = 0; i < s.length(); i++) {
      len = 0;
      unordered_map<char, int> m;
      for (int j = i; j < s.length(); j++) {
        if (m.find(s[j]) != m.end()) {
          break;
        } else {
          m.insert(make_pair(s[j], j));
          len++;
        }
      }
      maxLen = max(maxLen, len);
    }

    return maxLen;
  }
};
// @lc code=end
