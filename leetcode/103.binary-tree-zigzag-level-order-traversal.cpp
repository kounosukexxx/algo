/*
 * @lc app=leetcode id=103 lang=cpp
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */
#include <cmath>
#include <queue>
#include <string>
#include <vector>
using namespace std;

//  Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode *left;
  TreeNode *right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// 回し方が悪い
// reverse(begin, end)すれば配列逆にできる

// @lc code=start
class Solution {
public:
  vector<vector<int>> zigzagLevelOrder(TreeNode *root) {
    vector<vector<int>> ans;
    if (root == nullptr) {
      return ans;
    }

    vector<int> curr;
    queue<TreeNode *> q;
    int nodeCountInLevel = 0;
    int level = 1;
    bool goingRight = false;

    q.push(root);
    curr.emplace_back(root->val);
    ans.emplace_back(curr);
    curr.clear();
    while (!q.empty()) {
      TreeNode *popped = q.front();
      q.pop();

      if (popped->left != nullptr) {
        q.push(popped->left);
        curr.emplace_back(popped->left->val);
      }
      nodeCountInLevel++;
      if (popped->right != nullptr) {
        q.push(popped->right);
        curr.emplace_back(popped->right->val);
      }
      nodeCountInLevel++;

      if (nodeCountInLevel == pow(2, level)) {
        level++;
        if (!goingRight) {
          reverse(curr.begin(), curr.end());
        }
        if (!curr.empty()) {
          ans.emplace_back(curr);
          curr.clear();
        }
        nodeCountInLevel = 0;
        goingRight = !goingRight;
      }
    }

    return ans;
  }
};
// @lc code=end
