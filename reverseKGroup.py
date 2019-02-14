"""
https://leetcode.com/problems/reverse-nodes-in-k-group/
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
	def reverseKGroup(self, head, k):
		# a dummy node to ease the node manipulation with `head`
		dummy = ListNode(0)
		dummy.next = head
		
		pre = dummy  # maintain access to the current sub linked list
		current = head
		# count the number of nodes we have encountered
		count = 0
		
		while current:
			current = current.next
			count += 1
			# do the reverse operation only if we have counting enough node (k)
			if count % k == 0:
				# reverse the sub linked list of length k
				tail = pre.next
				for _ in range(k - 1):
					# insert node `tail.next` ahead the current `head` node
					head = pre.next
					pre.next = tail.next
					tail.next = pre.next.next
					pre.next.next = head
				# set pre to the last node of current sub linked list
				pre = tail
		return dummy.next