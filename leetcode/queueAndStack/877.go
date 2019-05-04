package leetcode

type MinStack struct {
	elems     []int
	min_elems []int
	top       int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{elems: []int{}, min_elems: []int{}, top: -1}
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)
	if this.top < 0 {
		this.min_elems = append(this.min_elems, x)
	} else {
		last_min := this.min_elems[this.top]
		if x < last_min {
			this.min_elems = append(this.min_elems, x)
		} else {
			this.min_elems = append(this.min_elems, last_min)
		}
	}
	this.top = this.top + 1
}

func (this *MinStack) Pop() {
	if this.top >= 0 {
		this.elems = this.elems[:this.top]
		this.min_elems = this.min_elems[:this.top]
		this.top = this.top - 1
	}
}

func (this *MinStack) Top() int {
	if this.top >= 0 {
		return this.elems[this.top]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if this.top >= 0 {
		return this.min_elems[this.top]
	}
	return 0
}
