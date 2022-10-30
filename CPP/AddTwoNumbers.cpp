/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *head = nullptr;
        ListNode *current = nullptr;
        int current_val = 0;
        int carry = 0;
        
        while(l1 != nullptr || l2 != nullptr || carry > 0) {
            if(l1 == nullptr && l2 == nullptr) {
                current_val = carry;
            } else if(l1 == nullptr && l2 != nullptr) {
                current_val = carry + l2->val;
                l2 = l2->next;
            } else if(l1 != nullptr && l2 == nullptr) {
                current_val = carry + l1->val;
                l1 = l1->next;
            } else {
                current_val = carry + l1->val + l2->val;
                l1 = l1->next;
                l2 = l2->next;
            }
            
            // if(current_val >= 10) {
            //     carry = 1;
            //     current_val = current_val - 10;
            // }
            carry = current_val / 10;
            current_val = current_val % 10;
            
            if(current == nullptr){
                head = new ListNode(current_val);
                current = head;
            } else {
                current->next = new ListNode(current_val);
                current = current->next;
            }
        }
        
        return head;
    }
};