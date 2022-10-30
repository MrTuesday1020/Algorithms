package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	studentTypeMap := map[int]int{
		0: 0,
		1: 0,
	}
	for i := range students {
		if students[i] == 0 {
			studentTypeMap[0] += 1
		} else {
			studentTypeMap[1] += 1
		}
	}
	
	for i := range sandwiches {
		if sandwiches[i] == 0 && studentTypeMap[0] > 0 {
			studentTypeMap[0] -= 1
		} else if sandwiches[i] == 1 && studentTypeMap[1] > 0 {
			studentTypeMap[1] -= 1
		} else {
			break
		}
	}
	return studentTypeMap[0] + studentTypeMap[1] 
}


func main() {
	fmt.Println(countStudents([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
	fmt.Println(countStudents([]int{1,0,1,0,1,1,0,1,1,1,1,0,0,0,1,1,1,0,1,1,1,1,0,0,0,1,0,0,0,0}, []int{0,1,0,0,1,1,1,1,1,1,0,1,1,0,0,0,1,1,0,0,1,1,1,1,0,0,1,0,1,0}))
}