package main

import "fmt"

var IntToRoman = []struct{
	digit int
	roman string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	ans := ""
	for _, m := range IntToRoman {
		for num >= m.digit {
			ans += m.roman
			num -= m.digit
		}
	}
	return ans
}

func main() {
	fmt.Println(intToRoman(22))
	fmt.Println(intToRoman(10))
}