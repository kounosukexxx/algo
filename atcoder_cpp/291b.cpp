#include <iostream>
#include <vector>
using namespace std;

int main() {
  int n;
  double sum = 0.0;
  cin >> n;

  vector<double> a(5 * n);
  for (int i = 0; i < a.capacity(); i++) {
    cin >> a[i];
  }

  for (int i = 0; i < a.capacity(); i++) {
    sort(a.begin(), a.end());
  }
  for (int i = n; i < 4 * n; i++) {
    sum += a[i];
  }
  cout << sum / ((double)(3 * n)) << endl;
}