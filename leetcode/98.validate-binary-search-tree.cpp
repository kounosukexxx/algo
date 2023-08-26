/*
 * @lc app=leetcode id=98 lang=cpp
 *
 * [98] Validate Binary Search Tree
 */

// Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right)
      : val(x), left(left), right(right) {}
};

// Brute Force: O(n^2)

// @lc code=start
class Solution {

public:
  bool existsInCorrectNode(TreeNode *root, int val) { // root: head
    if (root == nullptr) {
      return false;
    }
    if (root->val == val) {
      return true;
    }
    if (val < root->val) {
      return existsInCorrectNode(root->left, val);
    }
    if (root->val < val) {
      return existsInCorrectNode(root->right, val);
    }
    return false; // impossible
  }

  bool isValid(TreeNode *root, TreeNode *head, int parentVal) {
    if (head == nullptr) {
      return true;
    }
    if (root != head && parentVal == head->val) {
      return false;
    }
    if (root != head && head->val == root->val) {
      return false;
    }

    // if (head->left != nullptr && head->left->val >= head->val) {
    //   return false;
    // }
    // if (head->right != nullptr && head->right->val <= head->val) {
    //   return false;
    // }
    if (!existsInCorrectNode(root, head->val)) {
      return false;
    }

    return isValid(root, head->left, head->val) &&
           isValid(root, head->right, head->val);
  }

  bool isValidBST(TreeNode *root) { return isValid(root, root, root->val); }
};
// @lc code=end