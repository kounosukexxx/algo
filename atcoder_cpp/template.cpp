#include <algorithm>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;
typedef long long ll;

struct input {
  int n;
  input(int n) : n(n) {}
};

vector<int> run(input &in) {
  vector<int> ans;
  for (int i = 0; i < in.n; i++) {
    ans.emplace_back(i);
  }
  return ans;
}

int main() {
  // input
  int n;
  cin >> n;
  input *in = new input(n);

  // run
  vector<int> ans = run(*in);

  // output
  for (auto v : ans) {
    cout << v << endl;
  }

  return 0;
}