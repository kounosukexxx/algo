package l1676;

import java.util.*;

// NOTE: 通ってるけど、無駄ある。下から上に見ていくので、left + mid + right == sizeになった時点で終わりでいい

/**
 * Definition for a binary tree node.
 */
class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;

  TreeNode(int x) {
    val = x;
  }
}

class Solution {
  TreeNode ans = null;
  HashMap<Integer, TreeNode> nodes = new HashMap<>();

  public TreeNode lowestCommonAncestor(TreeNode root, TreeNode[] nodes) {
    // create nodes
    for (TreeNode node : nodes) {
      this.nodes.put(node.val, node);
    }
    getDescendantCount(root);
    return ans;
  }

  int getDescendantCount(TreeNode head) {
    if (head == null) {
      return 0;
    }

    int left = getDescendantCount(head.left);
    int mid = nodes.containsKey(head.val) ? 1 : 0;
    int right = getDescendantCount(head.right);

    switch (nodes.size()) {
      case 1:
        if (mid == 1) {
          setAns(head);
        }
        break;
      case 2:
        if ((left == 1 && mid == 1) || (mid == 1 && right == 1)) {
          setAns(head);
        } else if (left == 1 && right == 1) {
          setAns(head);
        }
        break;
      default:
        if ((left + mid + right == nodes.size()) && ((mid == 1) || (left > 0 && right > 0))) {
          setAns(head);
        }
        break;
    }

    return left + mid + right;

  }

  void setAns(TreeNode node) {
    if (ans == null) {
      ans = node;
    }
  }
}