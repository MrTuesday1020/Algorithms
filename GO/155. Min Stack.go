package main

type MinStack struct {
	stack []int
	min []int
}


func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min: []int{math.MaxInt32},
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack, val)
	currentMin := this.min[len(this.min)-1]
	if val < currentMin {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, currentMin)
	}
}


func (this *MinStack) Pop()  {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
		this.min = this.min[:len(this.min)-1]
	}
}


func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	
}