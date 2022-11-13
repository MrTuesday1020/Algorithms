package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums[]int, left, right, target int) int {
	location := partition(nums, left, right)
	if location == target {
		return nums[location]
	} else if location < target {
		return quickSelect(nums, location+1, right, target)
	} else {
		return quickSelect(nums, left, location-1, target)
	}
}

func partition(array []int, begin, end int) int {
	if begin == end {
		return begin
	}
	
	i := begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
	j := end       // array[end]是数组的最后一位

	// 没重合之前
	for i < j {
		if array[i] > array[begin] {
			array[i], array[j] = array[j], array[i] // 交换
			j--
		} else {
			i++
		}
	}

	/* 跳出while循环后，i = j。
	* 此时数组被分割成两个部分 -->  array[begin+1] ~ array[i-1] < array[begin]
	*                      -->  array[i+1] ~ array[end] > array[begin]
	* 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
	* 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
	*/
	// todo 这里没看明白 需要打印一下结果看下为啥！！！这里放在循环里是因为要判断=
	if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值组成时，会出现错误！
		i--
	}
	array[begin], array[i] = array[i], array[begin]
	return i
}

func main() {
	fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2))
}