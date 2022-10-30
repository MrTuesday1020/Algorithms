"""
input:
3 # the amount of numbers of list
1
3
4
2 # the amount of numbers of sublist
3
4
"""

class LinkedListNode:
	def __init__(self, x = None):
		self.val = x
		self.next = None

l = LinkedListNode()
subl = LinkedListNode()

m = int(input())
p = l
for i in range(m):
	l.next = LinkedListNode(int(input()))
	l = l.next
p = p.next
	
n = int(input())
q = subl
for i in range(n):
	subl.next = LinkedListNode(int(input()))
	subl = subl.next
q = q.next

def find(p, q):
	while p:
		if p.val == q.val:
			current1 = p
			current2 = q
			while current2:
				if current1.val == current2.val:
					if current2.next == None:
						return 1
					else:
						if current1.next == None:
							return -1
						else:
							current1 = current1.next
							current2 = current2.next
				else:
					return -1
		p = p.next
	return -1
		
print(find(p, q))