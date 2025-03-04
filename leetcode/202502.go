package leetcode

import (
	"Golang-algorithm/leetcode/2024"
	"container/heap"
	"fmt"
	"math"
	"slices"
	"sort"
)

func maxCount(m int, n int, ops [][]int) int {
	minRow := m
	minCol := n
	for i := 0; i < len(ops); i++ {
		row := ops[i][0]
		col := ops[i][1]
		minRow = min(minRow, row)
		minCol = min(minCol, col)
	}
	return minCol * minRow
}

func validPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return checkPalindrome(left+1, right, s) || checkPalindrome(left, right-1, s)
		}
		left++
		right--
	}
	return true
}

func checkPalindrome(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	//nums2 := make([]int, n)
	jIndex := 1
	oIndex := 0
	for oIndex < n && jIndex < n {
		if nums[jIndex]%2 == 1 {
			jIndex += 2
		} else if nums[oIndex]%2 == 0 {
			oIndex += 2
		} else {
			nums[jIndex], nums[oIndex] = nums[oIndex], nums[jIndex]
			jIndex += 2
			oIndex += 2
		}
	}
	return nums
}

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}
	var backtrack func(start int, cur []int)
	backtrack = func(start int, cur []int) {
		temp := make([]int, len(cur))
		copy(temp, cur)
		res = append(res, temp)
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			cur = append(cur, nums[i])
			backtrack(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0, []int{})
	return res
}

func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	n := len(nums)
	res := [][]int{}
	var backtrack func(cur []int)
	used := make([]bool, n)
	backtrack = func(cur []int) {
		if len(cur) == n {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 当出现重复元素时，比如输入 nums = [1,2,2',2'']，
			// 2' 只有在 2 已经被使用的情况下才会被选择，
			// 同理，2'' 只有在 2' 已经被使用的情况下才会被选择，这就保证了相同元素在排列中的相对位置保证固定。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, nums[i])
			backtrack(cur)
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	backtrack([]int{})
	return res
}

func removeDuplicates(nums []int) int {
	fast := 0
	slow := 0
	cnt := 0
	n := len(nums)
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && cnt < 3 {
			slow++
			nums[slow] = nums[fast]
		}
		cnt++
		fast++
		if fast < n && nums[fast] != nums[fast-1] {
			cnt = 0
		}
	}
	return slow + 1
}

func removeDuplicates_lin(nums []int) int {
	stackSize := 2 //数组的前两个元素默认保留
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[stackSize-2] {
			nums[stackSize] = nums[i]
			stackSize++
		}
	}
	return min(len(nums), stackSize)
}

type RangeFreqQuery struct {
	numIndexs map[int][]int
}

func Constructor0218(arr []int) RangeFreqQuery {
	inst := RangeFreqQuery{numIndexs: map[int][]int{}}
	for k, v := range arr {
		inst.numIndexs[v] = append(inst.numIndexs[v], k)
	}
	return inst
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	nums := this.numIndexs[value]
	start := sort.SearchInts(nums, left)
	// >right 相当于 >=right+1
	end := sort.SearchInts(nums, right+1)
	return end - start
}

func maxDistance(arrays [][]int) int {
	minNum := arrays[0][0]
	maxNum := arrays[0][len(arrays[0])-1]
	res := math.MinInt
	for i := 1; i < len(arrays); i++ {
		length := len(arrays[i])
		curMin := arrays[i][0]
		curMax := arrays[i][length-1]
		res = max(res, _024.Myabs(curMin, maxNum), _024.Myabs(curMax, minNum))
		minNum = min(minNum, curMin)
		maxNum = max(maxNum, curMax)
	}
	return res
}

func evenOddBit(n int) []int {
	//binaryList := []int{}
	even := 0
	odd := 0
	evenflag := true
	for n != 0 {
		temp := n % 2
		if evenflag {
			if temp == 1 {
				even++
			}
		} else {
			if temp == 1 {
				odd++
			}
		}
		evenflag = !evenflag
		//binaryList = append(binaryList, n%2)
		n = n / 2
	}
	return []int{even, odd}
}

func similarPairs(words []string) int {
	res := 0
	var encode func(word string) [26]int
	encode = func(word string) [26]int {
		wordEncode := [26]int{}
		for i := 0; i < len(word); i++ {
			wordEncode[word[i]-'a'] = 1
		}
		return wordEncode
	}
	wordCnt := make(map[[26]int]int)
	for i := 0; i < len(words); i++ {
		enCodeWord := encode(words[i])
		res += wordCnt[enCodeWord]
		wordCnt[enCodeWord]++
	}
	return res
}

type FoodRatings struct {
	foodMap    map[string]foodPair
	cuisineMap map[string]*foodHp
}

type foodHp []foodPair

func (f foodHp) Len() int {
	return len(f)
}

func (f foodHp) Less(i, j int) bool {
	a, b := f[i], f[j]
	if a.rating == b.rating {
		return a.foodName < b.foodName
	}
	return a.rating > b.rating
}

func (f foodHp) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *foodHp) Push(x any) {
	*f = append(*f, x.(foodPair))
}

func (f *foodHp) Pop() any {
	v := (*f)[len(*f)-1]
	*f = (*f)[:len(*f)-1]
	return v
}

type foodPair struct {
	rating   int
	cuisine  string
	foodName string
}

func ConstructorFood(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodMap := map[string]foodPair{}
	cuisineMap := map[string]*foodHp{}
	for i := 0; i < len(foods); i++ {
		rating, cuisine := ratings[i], cuisines[i]
		foodMap[foods[i]] = foodPair{
			rating:  rating,
			cuisine: cuisine,
		}
		if cuisineMap[cuisine] == nil {
			cuisineMap[cuisine] = &foodHp{}
		}
		heap.Push(cuisineMap[cuisine], foodPair{
			rating:   rating,
			foodName: foods[i],
		})
	}
	return FoodRatings{
		foodMap:    foodMap,
		cuisineMap: cuisineMap,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	p := this.foodMap[food]
	heap.Push(this.cuisineMap[p.cuisine], foodPair{
		rating:   newRating,
		foodName: food,
	})
	p.rating = newRating
	this.foodMap[food] = p
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.cuisineMap[cuisine]
	for h.Len() > 0 && (*h)[0].rating != this.foodMap[(*h)[0].foodName].rating {
		fmt.Println(heap.Pop(h))
	}
	return (*h)[0].foodName
}
