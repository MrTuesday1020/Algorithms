package main

import "fmt"

func prisonAfterNDays(cells []int, n int) []int {
	m := make(map[int]int)  // key cell value     value last position
	// try to find a loop
	newCells := cells
	distacne := 0
	for ; distacne < n; distacne++ {
		newCells = genNewCells(newCells)
		if distacne == n-1 {
			return newCells
		}
		num := CellsToInt(newCells)
		if _, ok := m[num]; ok {
			break
		}
		m[num] = distacne
	}
		
	n = n % distacne
	if n == 0 {
		n = distacne
	}
	
	for i := 0; i < n; i++ {
		cells = genNewCells(cells)
	}
	
	return cells
}

func genNewCells(cells []int) []int {
	ans := make([]int, 8)
	for j := 1; j < 7; j++ {
		if cells[j-1] == cells[j+1] {
			ans[j] = 1
		} else {
			ans[j] = 0
		}
	}
	return ans
}

func CellsToInt(cells []int) int {
	var ans int
	for i := 0; i < len(cells); i++ {
		ans = ans << 1;
		if (cells[i] == 1){
			ans = ans | cells[i];
		}
	}
	return ans
}

func IntToCells(num int) []int {
	var ans []int
	for num > 0 {
		ans = append(ans, num >> 1 & 1)
		num = num >> 1
	}
	return ans
}

func main() {
	fmt.Println(prisonAfterNDays([]int{1,0,0,1,0,0,1,0}, 1000000000))
	fmt.Println(IntToCells(10))
}