/*
 * @lc app=leetcode id=62 lang=cpp
 *
 * [62] Unique Paths
 */

// here: overflowしないように、割れる時に割っていく方式にした。

// reは自分でできてない↓
// recurse: 再帰でもできる。
// ⭐️後ろから求める:この答えを漸化式的に求めようとするとどうなるか？(editorial)
// 二次元配列: vector<vector<int>> dp(m, vector<int>(n, 1)) (1で埋まったn列配列を、m行もった配列)

// @lc code=start
class Solution {
public:
  int uniquePaths(int m, int n) {
    // see to it that m >= n
    if (m < n) {
      int tmp = m;
      m = n;
      n = tmp;
    }

    unsigned long long int res = 1;
    int numeratorElem = m + n - 2;
    int denominatorElem = n - 1;
    while (numeratorElem >= m) {
      res *= numeratorElem;
      if (res % denominatorElem == 0) {
        res /= denominatorElem;
        denominatorElem--;
      }
      numeratorElem--;
    }

    while (denominatorElem >= 1) {
      res /= denominatorElem;
      denominatorElem--;
    }
    return res;
  }
};
// @lc code=end
