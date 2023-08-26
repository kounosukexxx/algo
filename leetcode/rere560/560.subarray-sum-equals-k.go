package rere560

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

// loopをもっと再利用できる。

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums)+1)
	mp := make(map[int]*sumRecord, len(sums))

	sums[0] = 0
	mp[0] = &sumRecord{data: []int{0}, currIndex: 0}
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]

		_, ok := mp[sums[i+1]]
		if !ok {
			mp[sums[i+1]] = &sumRecord{data: []int{i + 1}, currIndex: 0}
		} else {
			mp[sums[i+1]].data = append(mp[sums[i+1]].data, i+1)
		}
	}

	ans := 0
	for i := 0; i < len(sums); i++ {
		mp[sums[i]].currIndex++
		target := sums[i] + k

		record, ok := mp[target]
		if !ok {
			continue
		}

		ans += len(record.data) - record.currIndex
	}

	return ans
}

type sumRecord struct {
	data      []int
	currIndex int
}

// @lc code=end
