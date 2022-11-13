package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	postCourses := make([][]int, numCourses)
	numOfPres := make([]int, numCourses)
	
	for _, prerequisite := range prerequisites {
		postCourses[prerequisite[1]] = append(postCourses[prerequisite[1]], prerequisite[0])
		numOfPres[prerequisite[0]] += 1
	}
	
	queue := []int{}
	for i := range numOfPres {
		if numOfPres[i] == 0 {
			queue = append(queue, i)
		}
	}
	
	result := []int{}
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		result = append(result, course)
		for _, c := range postCourses[course] {
			numOfPres[c] -= 1
			if numOfPres[c] == 0 {
				queue = append(queue, c)
			}
		}
	}
	
	if len(result) == numCourses {
		return result
	}
	return []int{}
}

func main() {
	
}