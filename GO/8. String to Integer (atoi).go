package main

func myAtoi(s string) int {
	negative := false
	num := 0
	
	// process prefix
	i := 0
	for i < len(s) {
		if s[i] == ' ' {
			i += 1
			continue
		} else if s[i] == '-' {
			negative = true
			i += 1
			break
		} else if s[i] == '+' {
			i += 1
			break
		} else {
			break
		}
	}
	
	// try to get a number
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num * 10 + int(s[i] - '0')
			if num > math.MaxInt32 {
				if negative {
					return math.MinInt32
				} else {
					return math.MaxInt32
				}
			}
		} else {
			break
		}
	}
	
	if negative {
		return -num
	}
	return num
}

func main() {
	
}