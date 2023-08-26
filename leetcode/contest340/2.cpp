#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
public:
  vector<long long> distance(vector<int> &nums) {
    unordered_map<int, vector<int>> mp; // nums's val -> index
    vector<int> iArray;
    for (int i = 0; i < nums.size(); i++) {
      // TODO: ここらへんはautoで綺麗に描ける
      int toEmplace;
      if (mp[nums[i]].size() == 0) {
        toEmplace = i;
      } else {
        toEmplace = mp[nums[i]][mp[nums[i]].size() - 1] + i;
      }
      mp[nums[i]].emplace_back(toEmplace);
      iArray.emplace_back(mp[nums[i]].size() - 1);
      // cout << nums[i] << toEmplace << mp[nums[i]].size() - 1 << endl;
    }

    vector<long long> ans;
    for (int i = 0; i < nums.size(); i++) {
      auto nx = mp.find(nums[i]);
      vector<int> n = nx->second;
      if (n.size() == 1) {
        ans.emplace_back(0);
        continue;
      }
      long long curr = 0;
      int I = iArray[i];
      curr += n[n.size() - 1] - n[i] - I * (n.size() - i);
      // cout << iArray[i] << " " << curr << endl;
      // cout << "a" << (n[n.size() - 1] - n[iArray[i]]) << " " << n[iArray[i]] * (n.size() - n[iArray[i]]) << endl;
      if (iArray[i] > 0) {
        curr -= n[i - 1] - I * i;
      }
      ans.emplace_back(curr);

      // for (int j = 0; j < nx.size(); j++) { // TODO:
      //   int curr = i - nx[j];
      //   if (curr < 0) {
      //     curr = -curr;
      //   }
      //   ans[i] += curr;
      // }
    }

    return ans;
  }
};

int main() {
  Solution s;
  vector<int> v;
  v.emplace_back(1);
  v.emplace_back(3);
  v.emplace_back(1);
  v.emplace_back(1);
  v.emplace_back(2);
  s.distance(v);
}
