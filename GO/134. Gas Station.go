package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	
	for start := 0; start < n; {
		sumGas, sumCost, step := 0, 0, 0
		for step < n {
			current := (start + step) % n
			sumGas += gas[current]
			sumCost += cost[current]
			if sumCost > sumGas {
				break
			}
			step += 1
		}
		if step == n {
			return start
		} else {
			start += step + 1
		}
	}
	
	return -1
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
}