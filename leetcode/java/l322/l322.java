package l322;
/*
 * @lc app=leetcode id=322 lang=java
 *
 * [322] Coin Change
 */

import java.util.*;

// @lc code=start
class Solution {
	public int coinChange(int[] coins, int amount) {
		if (amount == 0) {
			return 0;
		}

		ArrayList<Integer> dp = new ArrayList<>(Collections.nCopies(amount + 1, Integer.MAX_VALUE - 1));

		for (int j = 0; j < coins.length; j++) {
			if (coins[j] > amount) {
				continue;
			}
			dp.set(coins[j], 1); // nCopyしなかったらsetはerror。また、addはappendなので使えない
		}

		for (int i = 1; i < amount; i++) {
			for (int j = 0; j < coins.length; j++) {
				if ((long) i + (long) coins[j] > amount) {
					continue;
				}
				dp.set(i + coins[j], Math.min(dp.get(i + coins[j]), 1 + dp.get(i)));
			}
		}

		int ans = dp.get(amount);
		if (ans == Integer.MAX_VALUE - 1) {
			return -1;
		}
		return ans;
	}
}

public class l322 {
	public static void main(String[] args) {
		Solution s = new Solution();
		int[] coins = new int[] { 1, 2, 5 };
		int amount = 11;
	}
}
// @lc code=end
