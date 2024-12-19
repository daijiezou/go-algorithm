package top100

import (
	"slices"
	"sort"
	"strings"
)

func twoSum(nums []int, target int) []int {

	valToIndex := make(map[int]int)
	for k, v := range nums {
		valToIndex[v] = k
		if index, ok := valToIndex[target-v]; ok {
			return []int{k, index}
		}
	}
	return []int{}
}

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	myMap := make(map[string][]string)
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] > bytes[j]
		})
		myMap[string(bytes)] = append(myMap[string(bytes)], str)
	}
	for _, v := range myMap {
		res = append(res, v)
	}
	return res
}

func Encode(s string) string {
	cnts := make([]int, 26)
	for _, v := range s {
		cnts[v-'a']++
	}
	var res strings.Builder
	for i := 0; i < 26; i++ {
		res.WriteByte(byte(cnts[i]))
	}
	return res.String()
}

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)
	res := 0
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			res = max(res, cnt)
			cnt = 1
		}
	}
	res = max(res, cnt)
	return res
}
