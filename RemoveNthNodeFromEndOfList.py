"""
https://leetcode.com/problems/remove-nth-node-from-end-of-list/solution/
"""

# Definition for singly-linked list.
#class ListNode:
#	def __init__(self, x):
#		self.val = x
#		self.next = None

class Solution:
	def removeNthFromEnd(self, head, n):
		"""
		:type head: ListNode
		:type n: int
		:rtype: ListNode
		"""
		# invalid list
		if not head:
			return None
		
		p = head
		length = 0
		
		while p:
			length += 1
			p = p.next
		
		# remove first element	
		if length - n == 0:
			return head.next
		# invalid n
		if length - n < 0:
			return
		
		dummy = head
		cur = dummy
		for _ in range(length-n-1):
			cur = cur.next
		cur.next = cur.next.next
			
		return dummy