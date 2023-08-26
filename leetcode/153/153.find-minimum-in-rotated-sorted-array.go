package l153

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
func findMin(nums []int) int {
	leftMost := nums[0]
	l := 0
	r := len(nums) - 1
	if nums[l] <= nums[r] {
		return nums[l]
	}

	mid := (l + r) / 2
	for l != r {
		// TODO: 多分、r=lになったときだけループ抜けるんじゃなく、毎回midが最小かどうかチェックした方が早そう
		if leftMost <= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}
	return nums[mid]
}

// @lc code=end
