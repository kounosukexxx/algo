package l142

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func detectCycle(head *ListNode) *ListNode {
	for ; head != nil; head = head.Next {
		if head.Val == 1e6 {
			return head
		}
		head.Val = 1e6
	}
	return nil
}

// @lc code=end
