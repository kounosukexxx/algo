#include <algorithm>
#include <iostream>
#include <iterator> // std::back_inserter
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;
typedef long long ll;

struct T {
  int num;
  string str;
};

bool asc(const T &left, const T &right) { return left.num < right.num; }
bool ascDouble(const T &left, const T &right) { return left.num == right.num ? left.str < right.str : left.num < right.num; }

void print(vector<T> &array) {
  for (auto v : array) {
    cout << v.num << " " << v.str << endl;
  }
  cout << endl;
}

int main() {
  vector<T> array = {{1, "a"}, {2, "e"}, {1, "c"}, {3, "d"}, {2, "b"}};
  vector<T> array2;
  copy(array.begin(), array.end(), back_inserter(array2)); // array2にcopy

  // sort1
  sort(array.begin(), array.end(), asc);
  print(array);

  // sort2
  sort(array2.begin(), array2.end(), ascDouble);
  print(array2);

  return 0;
}