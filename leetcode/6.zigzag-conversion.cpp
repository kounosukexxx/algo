/*
 * @lc app=leetcode id=6 lang=cpp
 *
 * [6] Zigzag Conversion
 */

#include <string>
#include <unordered_map>
using namespace std;

// @lc code=start
class Solution {
public:
  string convert(string s, int numRows) {
    string res = "";
    int n = numRows - 1;
    for (int rowIndex = 0; rowIndex < numRows; rowIndex++) {
      int i = rowIndex;
      if (rowIndex == 0 || rowIndex == numRows - 1) {
        int indexToMove = n == 0 ? 1 : 2 * n;
        while (i < s.size()) {
          res += s[i];
          i += indexToMove;
        }
      } else {
        bool oddCount = true;
        int indexToMove = 2 * (n - rowIndex);
        while (i < s.size()) {
          res += s[i];
          i += indexToMove;

          oddCount = !oddCount;
          indexToMove = oddCount ? 2 * (n - rowIndex) : 2 * rowIndex;
        }
      }
    }

    return res;
  };
};

int main() {

}

// @lc code=end
