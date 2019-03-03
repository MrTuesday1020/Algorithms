"""
https://www.nowcoder.com/test/8537283/summary
P为给定的二维平面整数点集。定义 P 中某点x，如果x满足 P 中任意点都不在 x 的右上方区域内（横纵坐标都大于x），则称其为“最大的”。求出所有“最大的”点的集合。（所有点的横坐标和纵坐标都不重复, 坐标轴范围在[0, 1e9) 内）

输入描述:
第一行输入点集的个数 N， 接下来 N 行，每行两个数字代表点的 X 轴和 Y 轴。
对于 50%的数据,  1 <= N <= 10000;
对于 100%的数据, 1 <= N <= 500000;
输出描述:
输出“最大的” 点集合， 按照 X 轴从小到大的方式输出，每行两个数字分别代表点的 X 轴和 Y轴。

输入例子:
5
1 2
5 3
4 6
7 5
9 0
输出例子:
4 6
7 5
9 0
"""

n = int(input())

dictionary = {}

for i in range(n):
	numbers = input().split(" ")
	x = int(numbers[0])
	y = int(numbers[1])
	if x not in dictionary:
		dictionary[x] = y
	elif dictionary[x] < y:
		dictionary[x] = y

current_y = 0

ans = []

for key in reversed(sorted(dictionary)):
	if 	dictionary[key] >= current_y:
		ans.append((key, dictionary[key]))
		current_y = dictionary[key]
		
for point in reversed(ans):
	print(point[0], point[1])