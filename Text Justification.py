"""
https://leetcode.com/problems/text-justification/
"""

def fullJustify(words, maxWidth):
	ans = []
	currentLine = ""
	i = 0
	while i < len(words):
		if not currentLine:
			currentLine = words[i]
			# check if it is last line
			if i == len(words) - 1:
				# add extra space at the end 	
				ans.append(currentLine + " " * (maxWidth - len(currentLine)))
			# i+1 move to next word
			i += 1
		else:
			newLine = currentLine + " " + words[i]
			if len(newLine) == maxWidth:
				ans.append(newLine)
				# reset
				currentLine = ""
				# i+1 move to next word
				i += 1
			elif len(newLine) < maxWidth:
				currentLine = newLine
				# check if it is last line
				if i == len(words) - 1:
					ans.append(currentLine + " " * (maxWidth - len(currentLine)))
				# i+1 move to next word
				i += 1
			else:
				# there are enough words in the line, then adjust space
				# we don't do i+1 here because we the next word is longer than we need
				rest_space = maxWidth - len(currentLine)
				# if there is only one word(no space)
				if " " not in currentLine:
					currentLine += " " * rest_space		
					ans.append(currentLine)			
				else:
					# add space by replacing
					tmp = currentLine[0]
					j = 1
					while rest_space:
						if currentLine[j] == " " and currentLine[j-1] != " ":
							tmp += " " + currentLine[j]
							rest_space -= 1
						else:
							tmp += currentLine[j]
						
						# if there is some more space to be added
						if rest_space:
							# reach the end of current result
							if j == len(currentLine) - 1:
								j = 1
								currentLine, tmp = tmp, currentLine[0]
							else:
								j += 1
						else:
							currentLine = tmp[:-1] + currentLine[j:]
					ans.append(currentLine)
				# reset
				currentLine = ""
				
	return ans
		
words = ["Give","me","my","Romeo;","and,","when","he","shall","die,","Take","him","and","cut","him","out","in","little","stars,","And","he","will","make","the","face","of","heaven","so","fine","That","all","the","world","will","be","in","love","with","night","And","pay","no","worship","to","the","garish","sun."]
print(fullJustify(words, 25))