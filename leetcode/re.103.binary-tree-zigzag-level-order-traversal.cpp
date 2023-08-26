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

// ⭐️ある一定時間でのqueue.sizeを変数に保存する。len = q.size()
// - ⭐️一旦、whileの中の基本形（i使う部分）を模擬コードで書いてしまうのがいい

// level ごとにwhile, その中でwhile

// @lc code=start
class Solution {
public:
  vector<vector<int>> zigzagLevelOrder(TreeNode *root) {
    vector<vector<int>> ans;
    vector<int> curr;
    queue<TreeNode *> q;
    bool goRight = true;

    if (root == nullptr) {
      return ans;
    } else {
      q.push(root);
    }

    while (!q.empty()) {
      int size = q.size();
      while (size-- != 0) {
        TreeNode *front = q.front();
        q.pop();

        if (front->left != nullptr) {
          q.push(front->left);
        }
        if (front->right != nullptr) {
          q.push(front->right);
        }
        curr.emplace_back(front->val);
      }

      if (!goRight) {
        reverse(curr.begin(), curr.end());
      }

      goRight = !goRight;
      ans.emplace_back(curr);
      curr.clear();
    }

    return ans;
  }
};
// @lc code=end
