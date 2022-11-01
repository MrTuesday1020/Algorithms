package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	var meetings [][]int
	for i := range intervals {
		meetings = append(meetings, []int{intervals[i][0], 1})
		meetings = append(meetings, []int{intervals[i][1], -1})
	}
	
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] < meetings[j][0] {
			return true
		} else if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return false
	})
	
	cnt, ans := 0, 0
	for i := range meetings {
		cnt += meetings[i][1]
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(minMeetingRooms([][]int{
		[]int{0,30},
		[]int{5,10},
		[]int{15,20},
	}))
}