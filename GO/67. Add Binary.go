package main

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	l := la
	if lb > la {
		l = lb
	}
	
	ans := ""
	carry := 0
	for i := 0; i < l; i++ {
		if i < la {
			carry += int(a[la-i-1] - '0') // start computing from the end
		}
		if i < lb {
			carry += int(b[lb-i-1] - '0') // start computing from the end
		}
		current := carry % 2
		carry = carry / 2
		ans = fmt.Sprint(current) + ans
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}

func main() {
	
}