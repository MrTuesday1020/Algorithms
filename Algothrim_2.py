"""
	输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。要求不能创建任何新的结点，只能调整树中结点指针的指向。
	
	先中序遍历，将所有的节点保存到一个列表中。对这个list[:-1]进行遍历，每个节点的right设为下一个节点，下一个节点的left设为上一个节点。
"""


# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
	def Convert(self, pRootOfTree):
		# write code here
		if not pRootOfTree:
			return
		self.mid_array = []
		self.Mid(pRootOfTree)
		#for i,v in enumerate(self.mid_array[:-1]):
		#    v.right = self.mid_array[i + 1]
		#    self.mid_array[i + 1].left = v
		for i in range(len(self.mid_array)-1):
			self.mid_array[i].right = self.mid_array[i + 1]
			self.mid_array[i + 1].left = self.mid_array[i]
		return self.mid_array[0]
		
	def Mid(self,root):
		if not root:
			return
		self.Mid(root.left)
		self.mid_array.append(root)
		self.Mid(root.right)