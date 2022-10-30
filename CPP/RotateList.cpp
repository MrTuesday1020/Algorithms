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
    ListNode* rotateRight(ListNode* head, int k) {
        if(head == nullptr){
            return head;
        }
        
        int len = 1;
        ListNode* tmp = head;
        while(tmp->next){   //计算长度
            tmp = tmp->next;
            len++;
        }    
        k = k % len;
    
        tmp->next = head;   //首尾相连
        for(int i = 0; i < len - k; i++){   //旋转
            tmp = tmp->next;
        }
        
        ListNode* result = tmp->next;
        tmp->next = nullptr;    //切割
        return result; 
    }
};