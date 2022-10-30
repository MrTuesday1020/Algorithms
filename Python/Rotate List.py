"""
https://leetcode.com/problems/rotate-list/submissions/
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

def rotateRight(self, head: ListNode, k: int) -> ListNode:
	if not head:
		return None
	if not head.next:
		return head
	
	length = 1
	p = head
	while p.next:
		length += 1
		p = p.next
	# circle the link list into a ring    
	p.next = head
	# find the correct index to break the ring
	# index - 1 is new tail and index is new head
	index = length - k % length - 1
	p = head
	for i in range(index):
		p = p.next
	# find new head
	new_head = p.next
	# break the ring
	p.next = None
	return new_head