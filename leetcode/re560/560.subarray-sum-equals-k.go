package re560

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

// K*N spaceなので、out of memory

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	ans := 0
	dp := make(dp)
	for i := 1; i < len(sums); i++ {
		pair := newPair(i-1, sums[i]-k)
		dp.setDp(pair, sums)
		ans += dp[pair]
	}

	return ans
}

type pair struct {
	first  int
	second int
}

func newPair(first, second int) *pair {
	return &pair{
		first:  first,
		second: second,
	}
}

type dp map[*pair]int

func (dp dp) setDp(pair *pair, sums []int) {
	_, ok := dp[pair]
	if ok {
		return
	}

	if pair.first == 0 {
		if sums[0] == pair.second {
			dp[pair] = 1
		} else {
			dp[pair] = 0
		}
	} else {
		prePair := newPair(pair.first-1, pair.second)
		dp.setDp(prePair, sums)

		carry := 0
		if sums[pair.first] == pair.second {
			carry = 1
		}
		dp[pair] = dp[prePair] + carry
	}
}

// @lc code=end
