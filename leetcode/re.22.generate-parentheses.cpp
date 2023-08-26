/*
 * @lc app=leetcode id=22 lang=cpp
 *
 * [22] Generate Parentheses
 */
#include <string>
#include <vector>
using namespace std;

// NOTE: O(2^2*N)とかになることもある。⬅︎総当たりに近い時
// - 複雑かもってなったら、総当たりじゃないと無理かもしれない。複雑なやつはできないことある。
// recurse: 実験したら、i, jの時で一般的に試してみる。
// - この時は、木で表現できる。
// 単純にかっこが正しいのは、open>=closeを保ちつつ最後open=closeになる時
// - 他スタックで調べるのもできる（別のカッコが混じっててもいける）

// @lc code=start
class Solution {
public:
  vector<string> res;
  int n;

  void recurse(string curr, int open, int close) {
    if (open < close) {
      return; // invalid
    }

    if (curr.length() == 2 * n) {
      if (open == close) {
        res.push_back(curr); // valid
      }
      return; // valid or invalid
    }

    recurse(curr + "(", open + 1, close);
    recurse(curr + ")", open, close + 1);
  }

  vector<string> generateParenthesis(int n) {
    this->n = n;
    recurse("", 0, 0);
    return res;
  };
};
// @lc code=end
