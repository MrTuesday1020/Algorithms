package main

import "fmt"

func rotate(matrix [][]int)  {
	l := len(matrix)
	for depth := 0; depth < l / 2; depth++ {
		for i := depth; i < l - depth - 1; i++ {
			matrix[depth][i], matrix[i][l-depth-1] = matrix[i][l-depth-1], matrix[depth][i]
			matrix[depth][i], matrix[l-depth-1][l-i-1] = matrix[l-depth-1][l-i-1], matrix[depth][i]
			matrix[depth][i], matrix[l-i-1][depth] = matrix[l-i-1][depth], matrix[depth][i]
		} 
	} 
}


func main() {
	matrix := [][]int{
		[]int{5,1,9,11},
		[]int{2,4,8,10},
		[]int{13,3,6,7},
		[]int{15,14,12,16},
	}
	for _, l := range matrix {
		fmt.Println(l)
	}
	rotate(matrix)
	for _, l := range matrix {
		fmt.Println(l)
	}
}