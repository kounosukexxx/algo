#include <algorithm>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;
typedef long long ll;

int main() {
  // input
  int N, M;
  cin >> N >> M;
  vector<ll> A(N);
  ll allSum = 0;
  for (int i = 0; i < N; i++) {
    cin >> A[i];
    allSum += A[i];
  }
  sort(A.begin(), A.end());

  bool isFinishedOnce = false;
  ll x = A[0];
  ll isFirstValueZero = A[0] == 0;
  ll firstSum = A[0];
  ll currSum = A[0];
  ll maxSum = 0;
  for (int i = 1; i < N; i++) {
    if (A[i] == x || A[i] == x + 1) {
      currSum += A[i];
      if (!isFinishedOnce) {
        firstSum += A[i];
      }
    } else {
      isFinishedOnce = true;
      maxSum = max(maxSum, currSum);
      currSum = A[i];
    }
    x = A[i];
  }
  if (isFirstValueZero && isFinishedOnce && A[N - 1] == M - 1) {
    currSum += firstSum;
  }
  maxSum = max(maxSum, currSum);

  cout << allSum - maxSum << endl;
  return 0;
}