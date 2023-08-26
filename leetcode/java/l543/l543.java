package l543;
/*
 * @lc app=leetcode id=543 lang=java
 *
 * [543] Diameter of Binary Tree
 */

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}

// @lc code=start
/**
 * Definition for a binary tree node.
 */
class Solution {
    int max = 0;

    public int diameterOfBinaryTree(TreeNode root) {
        traversal(root);
        return max;
    }

    // return bigger depth
    int traversal(TreeNode head) {
        if (head == null) {
            return 0;
        }
        int leftDepth = traversal(head.left);
        int rightDepth = traversal(head.right);
        max = Math.max(max, leftDepth + rightDepth);
        return Math.max(leftDepth, rightDepth) + 1;
    }
}
// @lc code=end
