package main

import "fmt"

var RomanToInt = map[string]int{
	"I": 	1,
	"IV": 	4,
	"V": 	5,
	"IX": 	9,
	"X": 	10,
	"XL": 	40,
	"L": 	50,
	"XC":	90,
	"C": 	100,
	"CD": 	400,
	"D": 	500,
	"CM": 	900,
	"M": 	1000,
}

func romanToInt(s string) int {
	ans := 0
	i := 0
	for i < len(s) {
		if i + 1 < len(s) {
			roman := s[i:i+2]
			if digit, ok := RomanToInt[roman]; ok {
				ans += digit
				i += 2
				continue
			}
		}
		roman := string(s[i])
		digit := RomanToInt[roman]
		ans += digit
		i += 1
	}
	return ans
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}