package leetcode

import "sort"

// https://leetcode.cn/problems/reschedule-meetings-for-maximum-free-time-i/?envType=daily-question&envId=2025-07-09
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	frees := make([]int, n+1)
	frees[0] = startTime[0]
	for i := 1; i < n; i++ {
		frees[i] = startTime[i] - endTime[i-1]
	}
	frees[n] = eventTime - endTime[n-1]

	res := 0
	cur := 0
	left := 0
	for i := 0; i < n+1; i++ {
		cur += frees[i]
		if i < k {
			continue
		}
		res = max(res, cur)
		cur -= frees[left]
		left++
	}
	return res
}

func maxVowels(s string, k int) int {
	left := 0
	vowelMap := map[uint8]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	res := 0
	window := 0
	for i := 0; i < len(s); i++ {
		if vowelMap[s[k]] {
			window++
		}
		if i < k-1 {
			continue
		}
		if vowelMap[s[left]] {
			window--
		}
		left++
		res = max(res, window)
	}
	return res
}

// https://leetcode.cn/problems/reschedule-meetings-for-maximum-free-time-ii/?envType=daily-question&envId=2025-07-10
//func maxFreeTime2(eventTime int, startTime []int, endTime []int) int {
//
//}

// https://leetcode.cn/problems/count-days-without-meetings/?envType=daily-question&envId=2025-07-11
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	end := 0
	length := 0
	for _, x := range meetings {
		if x[0] > end { //不相交
			length += x[0] - end - 1
		}
		end = max(x[1], end)
	}
	length += days - end
	return length
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	n := len(trainers)
	j := 0
	for i := 0; i < len(players); i++ {
		for j < n && trainers[j] < players[i] {
			j++
		}
		if j == n {
			return i
		}
		// 匹配到一位训练师
		j++
	}
	return len(players)
}
