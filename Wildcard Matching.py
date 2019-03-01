def isMatch(s, p):
	i,j, match, idx = 0, 0, 0, -1
			
	while i < len(str(s)):
		if j < len(p) and (p[j] == '?' or s[i] == p[j]):
			i += 1
			j += 1
		# 碰到 *，假设它匹配空串，并且用 startIdx 记录 * 的位置，记录当前字符串的位置，p 后移
		elif j < len(p) and p[j] == '*':
			idx = j
			match = i
			j += 1
		# 当前字符不匹配，并且也没有 *，回退
		# p 回到 * 的下一个位置
		# match 更新到下一个位置
		# s 回到更新后的 match 
		# 这步代表用 * 匹配了一个字符
		elif idx != -1:
			j = idx + 1
			match += 1
			i = match
		else:
			return False
			
	while j < len(p) and p[j] == '*':
		j += 1;
		
	return j == len(p)
	
s = "aa"
p = "*"
print(isMatch(s, p))