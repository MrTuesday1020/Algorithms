package main

func maxSlidingWindow(nums []int, k int) []int {
	var ans, queue []int
	push := func(i int) {
		// 构造单调队列 后入队列的值必须小于前一个值
		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// keep the window size
		if i - queue[0] >= k {   // size = right - left + 1  所以这里要有=
			queue = queue[1:]
		}
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	ans = append(ans, nums[queue[0]])
	for i := k; i < len(nums); i++ {
		push(i)
		ans = append(ans, nums[queue[0]])
	}
	return ans
}

func main() {
	
}