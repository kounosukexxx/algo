#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
  bool isPrime(int v) {
    if (v < 2) {
      return false;
    }
    for (int p = 2; p * p <= v; p++) {
      if (v % p == 0)
        return false;
    }
    return true;
  }

public:
  int diagonalPrime(vector<vector<int>> &nums) {
    int ans = 0;
    for (int i = 0; i < nums[0].size(); i++) {
      int curr = nums[i][i];
      if (ans >= curr) {
        continue;
      }
      if (isPrime(curr)) {
        ans = curr;
      }
    }

    for (int i = 0; i < nums[0].size(); i++) {
      int curr = nums[i][nums[0].size() - 1 - i];
      if (ans >= curr) {
        continue;
      }
      if (isPrime(curr)) {
        ans = curr;
      }
    }
    return ans;
  }
};
