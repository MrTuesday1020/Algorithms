"""
https://leetcode.com/problems/merge-intervals/
https://leetcode.com/problems/insert-interval/
"""

# Definition for an interval.
# class Interval:
#     def __init__(self, s=0, e=0):
#         self.start = s
#         self.end = e

def merge(intervals):
	ans = []
	intervals = sorted(intervals, key=lambda i: i.start)
	for interval in intervals:
		if not ans:
			ans.append(interval)
		else:
			if ans[-1].end < interval.start:
				ans.append(interval)
			elif ans[-1].end <= interval.end:
				ans[-1].end = interval.end
	return ans
	
def insert(intervals, newInterval):
	if not intervals:
		return [newInterval]

	# 先找到插入位置
	intervals = sorted(intervals, key=lambda i: i.start)
	index = -1
	for i in range(len(intervals)):
		if newInterval.start <= intervals[i].start:
			index = i
			break
	if index == -1:
		index =len(intervals) - 1
		intervals = intervals + [newInterval]
	else:
		intervals = intervals[:index] + [newInterval] + intervals[index:]
		
	# 再合并
	ans = intervals[:index]
	for i in range(index, len(intervals)):
		if not ans:
			ans.append(intervals[i])
		else:
			if ans[-1].end < intervals[i].start:
				ans.append(intervals[i])
			elif ans[-1].end <= intervals[i].end:
				ans[-1].end = intervals[i].end
	return ans
	
	