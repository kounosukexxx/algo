package l283

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	mostLeftZero := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if mostLeftZero == -1 {
				mostLeftZero = i
			}
			continue
		}

		if mostLeftZero == -1 {
			continue
		}
		nums[mostLeftZero], nums[i] = nums[i], nums[mostLeftZero]
		mostLeftZero++
	}
}

// @lc code=end
