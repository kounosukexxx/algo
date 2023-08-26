package l236;
/*
 * @lc app=leetcode id=236 lang=java
 *
 * [236] Lowest Common Ancestor of a Binary Tree
 */

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

// @lc code=start
/**
 * Definition for a binary tree node.
 */
class Solution {
    boolean pFound, qFound;
    TreeNode ans;

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        findPorQ(root, p, q);
        return ans;
    }

    boolean findPorQ(TreeNode head, TreeNode p, TreeNode q) {
        if (head == null) {
            return false;
        }

        if (head.val == p.val) {
            pFound = true;
            if (!qFound) {
                if (findPorQ(head.left, p, q)) {
                    ans = head;
                } else if (findPorQ(head.right, p, q)) {
                    ans = head;
                }
            }
            return true;
        } else if (head.val == q.val) {
            qFound = true;
            if (!pFound) {
                if (findPorQ(head.left, p, q)) {
                    ans = head;
                } else if (findPorQ(head.right, p, q)) {
                    ans = head;
                }
            }
            return true;
        }

        boolean foundInLeft = findPorQ(head.left, p, q);
        boolean foundInRight = findPorQ(head.right, p, q);
        if (foundInLeft && foundInRight) {
            ans = head;
            return true;
        } else if (foundInLeft || foundInRight) {
            return true;
        } else {
            return false;
        }
    }
}
// @lc code=end
