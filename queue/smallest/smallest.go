package smallest

import "math"

type SmallestInfiniteSet struct {
	List   []int
	numMap map[int]struct{}
}

func Constructor() SmallestInfiniteSet {
	myList := make([]int, 0, 1000)
	myMap := make(map[int]struct{}, 1000)
	for i := 1; i <= 1000; i++ {
		myList = append(myList, i)
		myMap[i] = struct{}{}
	}
	return SmallestInfiniteSet{
		List:   myList,
		numMap: myMap,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	minNum := math.MaxInt32
	minIndex := 0
	length := len(this.List)
	for i := 0; i < length; i++ {
		if this.List[i] < minNum {
			minNum = this.List[i]
			minIndex = i
		}
	}
	popNum := this.List[minIndex]
	for i := minIndex; i < length-1; i++ {
		this.List[i] = this.List[i+1]
	}
	this.List = this.List[:length-1]
	delete(this.numMap, popNum)
	return popNum
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.numMap[num]; ok {
		return
	}
	this.List = append(this.List, num)
	this.numMap[num] = struct{}{}
}
