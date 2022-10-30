"""
https://leetcode.com/problems/merge-k-sorted-lists/
"""

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
	def mergeKLists(self, lists):
		"""
		:type lists: List[ListNode]
		:rtype: ListNode
		"""
		if not lists:
			return
		
		def mergeTwoLists(l1, l2):
			"""
			:type l1: ListNode
			:type l2: ListNode
			:rtype: ListNode
			"""
			if not l1:
				return l2
			if not l2:
				return l1

			head = ListNode(None)
			current = head

			while l1 and l2:
				if l1.val <= l2.val:
					current.next = l1
					l1 = l1.next
				else:
					current.next = l2
					l2 = l2.next
				current = current.next

			if not l1:
				current.next = l2
			if not l2:
				current.next = l1

			return head.next
		
		def DivideAndConquer(lists):
			#Merge with Divide And Conquer
			amount = len(lists)
			interval = 1
			while interval < amount:
				for i in range(0, amount - interval, interval * 2):
					lists[i] = mergeTwoLists(lists[i], lists[i + interval])
				interval *= 2
			return lists[0]
		
		def OneByOne(lists):
			#Merge one by one
			for i in range(1,len(lists)):
				lists[0] = mergeTwoLists(lists[0], lists[i])
			return lists[0]
		
		return DivideAndConquer(lists)