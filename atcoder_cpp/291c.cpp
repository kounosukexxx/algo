#include <iostream>
#include <set>
using namespace std;

int main() {
  int n;
  cin >> n;
  string str;
  cin >> str;
  set<pair<int, int>> used;
  used.insert(make_pair(0, 0));

  int x = 0, y = 0;
  for (auto c : str) {
    if (c == 'R')
      x++;
    else if (c == 'L')
      x--;
    else if (c == 'U')
      y++;
    else if (c == 'D')
      y--;

    if (used.find(make_pair(x, y)) != used.end()) {
      cout << "Yes" << endl;
      return 0;
    }

    used.insert(make_pair(x, y));
  }

  cout << "No" << endl;
  return 0;
}