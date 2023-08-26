package l112

/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// NOTE: intは負の場合もある。
// root.valが負だと、targetSum-root.valはtargetSumよりおおきくなる.


// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func hasPathSum(root *TreeNode, targetSum int) bool {
	switch {
	case root == nil:
		return false
	case targetSum == root.Val && root.Left == nil && root.Right == nil:
		return true
	// ここはイラン。targetSumは負の場合もある
	// case targetSum <= root.Val:
	// 	return false
	default:
		return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
	}
}

// @lc code=end
