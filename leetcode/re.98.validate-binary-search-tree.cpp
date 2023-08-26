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
  TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// upperLimit and lowerLimit:
// - O(N), O(N)
// - DFS
//   - recursion ⬅︎
//   - stackもいける (recursionと同じく、調べる順にpushしていく。そんでwhileでやる)
//   - AND in-order traversal: subtree単位でみると、葉とparentの値だけでvalidateできる。
//     - elements show in ascending order
//     - stackもいける。右のやつらがどんどんpushされてく。先によまれる順番を理解するとはやそう。

// @lc code=start
class Solution {
  bool isValid(TreeNode *head, int *upperLimit, int *lowerLimit) {
    if (head == nullptr) {
      return true;
    }
    if (upperLimit != nullptr && head->val >= *upperLimit) {
      return false;
    }
    if (lowerLimit != nullptr && head->val <= *lowerLimit) {
      return false;
    }

    return isValid(head->left, &head->val, lowerLimit) && isValid(head->right, upperLimit, &head->val);
  }

public:
  bool isValidBST(TreeNode *root) { return isValid(root, nullptr, nullptr); }
};
// @lc code=end