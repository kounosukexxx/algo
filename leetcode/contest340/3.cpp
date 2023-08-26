#include <algorithm>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
public:
  int minimizeMax(vector<int> &nums, int p) {
    sort(nums.begin(), nums.end());
    return nums[2 * p - 1] - nums[2 * p - 2];
  }
};