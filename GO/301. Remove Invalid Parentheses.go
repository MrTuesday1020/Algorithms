package main

func removeInvalidParentheses(s string) []string {
	leftRemoval, rightRemoval := 0, 0
	for i := range s {
		if s[i] == '(' {
			leftRemoval += 1
		} else if s[i] == ')' {
			if leftRemoval == 0 {
				rightRemoval += 1
			} else {
				leftRemoval -= 1
			}
		}
	}
	
	ans := []string{}
	helper(&ans, s, 0, leftRemoval, rightRemoval)
	return ans
}

func helper(ans *[]string, str string, start, leftRemoval, rightRemoval int) {
	if leftRemoval == 0 && rightRemoval == 0 {
		if valid(str) {
			*ans = append(*ans, str)
		}
		return
	}
	
	for i := start; i < len(str); i++ {
		//为了使结果去重，也就是”(((((()“情况，重复的左括号去掉任意一个就行了
		if i > start && str[i] == str[i-1] {
			continue
		}
		// 如果剩余的字符无法满足去掉的数量要求，直接返回
		if leftRemoval + rightRemoval > len(str) - i {
			return
		}
		// 尝试去掉一个左括号
		if leftRemoval > 0 && str[i] == '(' {
			helper(ans, str[:i]+str[i+1:], i, leftRemoval-1, rightRemoval)
		}
		// 尝试去掉一个右括号
		if rightRemoval > 0 && str[i] == ')' {
			helper(ans, str[:i]+str[i+1:], i, leftRemoval, rightRemoval-1)
		}
	}
}

func valid(str string) bool {
	count := 0
	for i := range str {
		if str[i] == '(' {
			count += 1
		} else if str[i] == ')' {
			if count  == 0 {
				return false
			} else {
				count -= 1
			}
		}
	}
	return count == 0
}

func main() {
	
}