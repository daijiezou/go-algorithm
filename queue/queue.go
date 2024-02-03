package queue

type FrontMiddleBackQueue struct {
	List   []int
	length int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.List = append(this.List, 0)
	for i := this.length; i > 0; i-- {
		this.List[i] = this.List[i-1]
	}
	this.List[0] = val
	this.length++
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	this.List = append(this.List, 0)
	mid := this.length / 2
	for i := this.length; i > mid; i-- {
		this.List[i] = this.List[i-1]
	}
	this.List[mid] = val
	this.length++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.List = append(this.List, val)
	this.length++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.length == 0 {
		return -1
	}
	popNum := this.List[0]
	for i := 0; i < this.length-1; i++ {
		this.List[i] = this.List[i+1]
	}
	this.length--
	this.List = this.List[:this.length]
	return popNum
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.length == 0 {
		return -1
	}
	var mid int
	if this.length%2 == 0 {
		mid = (this.length - 1) / 2
	} else {
		mid = this.length / 2
	}

	popNum := this.List[mid]
	for i := mid; i < this.length-1; i++ {
		this.List[i] = this.List[i+1]
	}
	this.length--
	this.List = this.List[:this.length]
	return popNum
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.length == 0 {
		return -1
	}
	popNum := this.List[this.length-1]
	this.length--
	this.List = this.List[:this.length]
	return popNum
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
