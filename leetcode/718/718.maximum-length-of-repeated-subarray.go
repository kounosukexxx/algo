package l718

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// ex2でtimeoutするだろう

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	A := nums1
	B := nums2
	if len(A) > len(B) {
		A, B = B, A
	}

	// int: B val ➡︎ []int: indexes
	bMap := make(map[int][]int, len(B))
	for i, v := range B {
		bMap[v] = append(bMap[v], i)
	}

	ans := 0
	currSubarraies := make([][]int, 0, len(B))
	indexes := bMap[A[0]]
	for _, index := range indexes {
		currSubarraies = append(currSubarraies, []int{index})
	}

	for i := 1; i < len(A); i++ {
		tmpSubarraies := make([][]int, len(currSubarraies))
		for j := 0; j < len(currSubarraies); j++ {
			tmpSubarraies[j] = make([]int, len(currSubarraies[j]))
			copy(tmpSubarraies[j], currSubarraies[j])
		}

		a := A[i]
		for _, subArray := range tmpSubarraies {
			
			if B[subArray[len(subArray)-1]] == a {
				
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
