#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>
using namespace std;

// NOTE: とにかく、draftを忠実に再現するのが大事。コードをみて考えるとかなり難しい。

struct input {
  unordered_map<string, string> mp;
  input(unordered_map<string, string> mp) : mp(mp) {}
};

bool cycles(unordered_map<string, string> &mp, unordered_set<string> &visited, string head) {
  if (visited.find(head) != visited.end()) {
    return false;
  }
  visited.emplace(head);

  for (auto itr = mp.find(head); itr != mp.end(); itr = mp.find(itr->second)) {
    // cycles
    if (itr->second == head) {
      return true;
    }

    if (visited.find(itr->second) != visited.end()) {
      return false;
    }

    visited.emplace(itr->second);
  }
  return false;
}

string run(input &in) {
  unordered_set<string> visited;
  for (auto v : in.mp) {
    if (cycles(in.mp, visited, v.first)) {
      return "No";
    }
  }
  return "Yes";
}

int main() {
  // input
  int n;
  unordered_map<string, string> mp;
  cin >> n;
  for (int i = 0; i < n; i++) {
    string s, t;
    cin >> s >> t;
    mp.emplace(s, t);
  }
  input *in = new input(mp);

  // run
  string ans = run(*in);

  // output
  cout << ans << endl;

  return 0;
}