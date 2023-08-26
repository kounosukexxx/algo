#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

// nが感覚とずれているの注意
// dp[n][k][d] := n番目までの数を使って、k個の数を選んで、その和をdで割った余りが0になるように選んだ時の、その和の最大値
int main() {
  // input
  int N, K, D;
  cin >> N >> K >> D;

  vector<vector<vector<long long>>> dp(N + 1, vector<vector<long long>>(K + 1, vector<long long>(D, -1)));
  dp[0][0][0] = 0;
  for (int n = 0; n < N; n++) {
    int a;
    cin >> a;
    for (int k = 0; k < K + 1; k++) { // k=Kの時、任意のnの時のdpを埋めたいのでkまで。
      for (int d = 0; d < D; d++) {
        if (dp[n][k][d] == -1) {
          continue;
        }

        dp[n + 1][k][d] = max(dp[n + 1][k][d], dp[n][k][d]);

        if (k == K) {
          continue;
        }

        dp[n + 1][k + 1][(d + a) % D] = max(dp[n + 1][k + 1][(d + a) % D], dp[n][k][d] + a);
      }
    }
  }
  cout << dp[N][K][0] << endl;

  return 0;
}