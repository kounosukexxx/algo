#include <iostream>
#include <vector>
using namespace std;
bool satisfies();

struct card {
  int front, back;
};

int main() {
  int n;
  cin >> n;

  vector<card> cards(n);
  for (int i = 0; i < 2 * n; i++) {
    if (i % 2 == 0) {
      cin >> cards[i / 2].front;
    } else {
      cin >> cards[i / 2].back;
    }
  }

  // print
  // for (int i = 0; i < n; i++) {
  //   cout << cards.at(i).front << ", ";
  //   cout << cards.at(i).back << "\n";
  // }
  // cout << endl;

  int ans = 0;
  for (int i = 0; i < n - 1; i++) {
  }
  return 0;
}
