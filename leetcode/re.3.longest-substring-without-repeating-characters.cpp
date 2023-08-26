/*
 * @lc app=leetcode id=3 lang=cp
 *
 * [3] Longest Substring Without Repeating Characters
 */

#include <string>
#include <unordered_map>
using namespace std;

// NOTE:
// mapは一回だけ作るようにする。brute-forceでやらない。つまり、一個ずつ触らない。
// right&leftなど、indexを指すvarを定義する
// 記憶のためにmapを仕様するのはok
// `found`, `not found`って書くのいい。
// []のネストを間違えないように注意
// while about
// Xの時、Xはinvalidに最後なりうる。（ってことはincrement最後でやるのがいいかも）
// 共通のもののみifにいれる。
// @lc code=start
class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    int maxLen = 0;
    unordered_map<char, int> exist;
    int left = 0, right = 0;
    while (right < s.length()) {
      if (exist.find(s[right]) == exist.end() || exist[s[right]] == 0) {
        // not found
      } else {
        // found
        while (s[left] != s[right]) {
          exist[s[left]]--;
          left++;
        }
        exist[s[left]]--;
        left++;
      }

      exist[s[right]]++;
      maxLen = max(maxLen, right - left + 1);
      right++; // 最後
    }
    return maxLen;
  };
};
// @lc code=end
// abcdbdabc
// " "
// tmmzuxt