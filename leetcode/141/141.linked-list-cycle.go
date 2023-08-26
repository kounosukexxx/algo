package l141

/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func hasCycle(head *ListNode) bool {
	for ; head != nil; head = head.Next {
		if head.Val == 1e6 {
			return true
		}
		head.Val = 1e6
	}
	return false
}

// @lc code=end
