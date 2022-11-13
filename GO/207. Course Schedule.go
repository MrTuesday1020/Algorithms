package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	postCourses := make([][]int, numCourses)	// 记录课程i对应的是哪些课程的前置课程
	numOfPres := make([]int, numCourses)		// 记录课程i所依赖的前置可能有几个
	
	for _, prerequisite := range prerequisites {
		postCourses[prerequisite[1]] = append(postCourses[prerequisite[1]], prerequisite[0])
		numOfPres[prerequisite[0]] += 1
	}
	
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if numOfPres[i] == 0 {
			queue = append(queue, i)   // 入度为0表示不依赖任何其他的课程
		}
	}
	
	result := []int{}
	for len(queue) > 0 {
		course := queue[0]		// 取出一门入度为0的课（即没有前置课程的课）
		queue = queue[1:]
		result = append(result, course)
		for _, c := range postCourses[course] {
			// 便利当前课的所有后置课程，并将入度减一（当前课修完了，前置课程少了一门了）
			numOfPres[c] -= 1
			if numOfPres[c] == 0 {	// 如果该后置课程的入度为0了，表示没有前置课程了，可以学了
				queue = append(queue, c)
			}
		}
	}
	
	return len(result) == numCourses
}

func main() {
	
}