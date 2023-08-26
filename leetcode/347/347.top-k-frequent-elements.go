package l347

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int, len(nums))
	for _, num := range nums {
		mp[num]++
	}
	pairs := make([]*pair, 0, len(mp))
	for k, v := range mp {
		pairs = append(pairs, &pair{num: k, frequency: v})
	}

	// sort pairs in descending order of frequency
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].frequency > pairs[j].frequency
	})

	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, pairs[i].num)
	}

	return ans
}

type pair struct {
	num       int
	frequency int
}
