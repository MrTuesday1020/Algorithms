package main

type st struct {
	num int
	cnt int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]*st)
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]].cnt += 1
		} else {
			m[nums[i]] = &st{
				num: nums[i],
				cnt: 1,
			}
		}
	}
	
	arr := []*st{}
	for _, v := range m {
		arr = append(arr, v)
	}
	
	sort.Slice(arr, func(i, j int) bool {  // todo 这里自己实现快排的话 复杂度会降低 因为只要topK
		return arr[i].cnt > arr[j].cnt
	})
	
	
	ans := []int{}
	for i := 0; i < k && i < len(arr); i++ {
		ans = append(ans, arr[i].num)
	}
	return ans
}

func main() {
	
}