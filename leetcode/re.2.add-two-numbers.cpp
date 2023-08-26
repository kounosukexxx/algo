/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 */

// NOTE: define dummyHead, and answer is dummyHead->next
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
    ListNode *dummyHead = new ListNode();
    ListNode *cur = dummyHead;
    int carry = 0;
    int val1;
    int val2;

    while (canContinue(l1, l2, carry)) {
      val1 = l1 ? l1->val : 0;
      val2 = l2 ? l2->val : 0;
      int sum = val1 + val2 + carry;

      cur->next = new ListNode(sum % 10, nullptr);
      cur = cur->next;

      carry = sum / 10;

      l1 = l1 ? l1->next : nullptr;
      l2 = l2 ? l2->next : nullptr;
    }

    return dummyHead->next;
  }
};
// @lc code=end
