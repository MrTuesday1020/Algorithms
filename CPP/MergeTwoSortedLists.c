/*
https://leetcode.com/problems/merge-two-sorted-lists/
*/

#include <stdio.h>

/**
 * Definition for singly-linked list.
*/
struct ListNode {
	int val;
	struct ListNode *next;
};

struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2) {
	if(!l1)
		return l2;
	if(!l2)
		return l1;
	
	struct ListNode* head = NULL;
	struct ListNode* current = NULL;
	while (l1 && l2) {
		struct ListNode *temp;
		if (l1->val <= l2->val) {
			temp = l1;
			l1 = l1->next;
		} else {
			temp = l2;
			l2 = l2->next;
		}
		
		// initial head
		if (head == NULL) {
			head = current = temp;
		} else {
			current->next = temp;
			current = temp;
		}
	}
	
	if(!l1)
		current->next = l2;
	if(!l2)
		current->next = l1;
	
	return head;
}

int main(int argc, char *argv[]) {

}