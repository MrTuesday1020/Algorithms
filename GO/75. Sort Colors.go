package main

func sortColors(nums []int)  {
	var red []int
	var white []int
	var blue []int
	
	for i := range nums {
		if nums[i] == 0 {
			red = append(red, 0)
		} else if nums[i] == 1 {
			white = append(white, 1)
		} else if nums[i] == 2 {
			blue = append(blue, 2)
		}
	}
	red = append(red, white...)
	red = append(red, blue...)
	copy(nums, red)
}

func sortColors(nums []int)  {
	p0 := 0 			// means the position of next 0
	p1 := len(nums)-1	// means the position of next 2
	
	for i := 0; i <= p1; {
		if nums[i] == 0 {
			// num == 0, swap with p0, move index and p0 forward
			nums[i], nums[p0] = nums[p0], nums[i]
			p0 += 1
			i +=1
		} else if nums[i] == 1 {
			// num == 0, it is located where it should be, move index forward only
			i +=1
		} else {
			// num == 2, swap with p2, move p1 backward
			// because we do not know whether num[p1](before swap) == 1 or 2, so we not move index
			nums[i], nums[p1] = nums[p1], nums[i]
			p1 -= 1
		}
	}
}

func main() {
	
}