package re322;
/*
 * @lc app=leetcode id=322 lang=java
 *
 * [322] Coin Change
 */

// TODO: topdown

import java.util.*;

// @lc code=start
class Solution {
  // Intger.MAX_VALUE means unreachable
  public int coinChange(int[] coins, int amount) {
    int[] dp = new int[1 + amount];
    int ans = setDp(dp, coins, amount);
    if (ans == Integer.MAX_VALUE) {
      ans = -1;
    }
    return ans;
  }

  int setDp(int[] dp, int[] coins, int i) {
    if (i == 0) {
      return 0;
    }
    if (dp[i] != 0) {
      return dp[i];
    }

    int min = Integer.MAX_VALUE - 1;
    for (int coin : coins) {
      if (i - coin < 0) {
        continue;
      }

      min = Math.min(min, setDp(dp, coins, i - coin));
    }

    dp[i] = min + 1;
    return dp[i];
  }
}

public class re322 {
  public static void main(String[] args) {
    Solution s = new Solution();
    int[] coins = new int[] { 1, 2, 5 };
    int amount = 11;
    System.out.println(s.coinChange(coins, amount));
  }
}
// @lc code=end
