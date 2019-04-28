"""
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

def deleteDuplicates(self, head: ListNode) -> ListNode:
	if not head:
		return
	
	fakehead = ListNode(0)
	fakehead.next = head
	prev = fakehead
	current = head
	
	while current:
		while current.next and current.val == current.next.val:
			current = current.next
		# no duplicates
		if prev.next == current:
			prev = prev.next
		# duplicates
		else:
			prev.next = current.next
		
		current = current.next
	
	return fakehead.next