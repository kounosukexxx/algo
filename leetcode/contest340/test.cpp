#include <string>
#include <unordered_map>
#include <vector>
#include <iostream>
using namespace std;

int main() {
  unordered_map<int, vector<int> > mp; // nums's val -> index
  mp[1].push_back(1);
  mp[1].push_back(2);
  mp[2].push_back(3);
  cout << mp[1][0] << mp[1][1] << mp[2][0] << endl;
  return 0;
}