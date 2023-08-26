#include <iostream>
#include <stack>
#include <string>
#include <unordered_set>
#include <vector>
using namespace std;

int main() {
  // input
  string str;
  cin >> str;
  stack<char> st;
  unordered_set<char> set;

  for (auto c : str) {
    switch (c) {
    case '(': {
      st.push('(');
      break;
    }
    case ')': {
      char top = st.top();
      while (top != '(') {
        set.erase(top);
        st.pop();
        top = st.top();
      }
      st.pop(); // pop (
      break;
    }
    default:
      if (set.find(c) != set.end()) {
        cout << "No" << endl;
        return 0;
      }
      set.insert(c);
      st.push(c);
      break;
    }
  }

  cout << "Yes" << endl;
  return 0;
}