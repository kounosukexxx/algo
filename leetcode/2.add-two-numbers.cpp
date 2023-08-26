/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 */

// NOTE: 重複コードは書かない！
// LEARN: 長い数字は、linked listで表現する手がある。
//  Definition for singly-linked list.
struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// @lc code=start
class Solution {
  bool canContinue(ListNode *l1, ListNode *l2, int carry) {
    return (l1 != nullptr || l2 != nullptr) || (carry > 0);
  }

public:
  ListNode *addTwoNumbers(ListNode *l1, ListNode *l2) {
    ListNode *ans = new ListNode();
    ListNode *tmp = ans;
    int carry = 0;
    int val1;
    int val2;

    while (canContinue(l1, l2, carry)) {
      if (l1 != nullptr) {
        val1 = l1->val;
      } else {
        val1 = 0;
      }
      if (l2 != nullptr) {
        val2 = l2->val;
      } else {
        val2 = 0;
      }

      int sum = val1 + val2 + carry;
      tmp->val = sum % 10;
      carry = sum / 10;

      if (l1 != nullptr)
        l1 = l1->next;
      if (l2 != nullptr)
        l2 = l2->next;

      if (canContinue(l1, l2, carry)) {
        tmp->next = new ListNode(0, nullptr);
        tmp = tmp->next;
      }
    }

    return ans;
  }
};
// @lc code=end
