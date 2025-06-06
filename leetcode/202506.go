package leetcode

// https://leetcode.cn/problems/distribute-candies-among-children-ii/?envType=daily-question&envId=2025-06-01
/*
给你两个正整数 n 和 limit 。
请你将 n 颗糖果分给 3 位小朋友，确保没有任何小朋友得到超过 limit 颗糖果，
请你返回满足此条件下的 总方案数 。
*/
func distributeCandies(n int, limit int) int64 {
	ans := int64(0)
	for i := 0; i <= min(n, limit); i++ {

		// 剩下的糖果超过2*limit,不满足条件
		if n-i > 2*limit {
			// 不存在合法的方案
			continue
		}
		// 第二个小孩最多可以分
		secondMax := min(limit, n-i)

		// 第二个小孩最少得分
		secondMin := max(0, n-i-limit)
		ans += int64(secondMax - secondMin + 1)
	}
	return ans
}

func candy(ratings []int) int {
	n := len(ratings)
	ans := 0
	for i := 0; i < n; i++ {
		start := i
		if i > 0 && ratings[i-1] < ratings[i] {
			start--
		}
		for i+1 < n && ratings[i+1] > ratings[i] {
			i++
		}
		top := i
		for i+1 < n && ratings[i] > ratings[i+1] {
			i++
		}
		inc := top - start
		desc := i - top
		ans += max(inc, desc) + ((inc-1)*inc+(desc-1)*desc)/2
	}
	return ans + n
}

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	total := 0
	n := len(status)
	visited := make([]int, n)
	var bfs func(boxs []int)
	bfs = func(boxs []int) {
		for len(boxs) > 0 {
			length := len(boxs)
			for _, box := range boxs {
				if status[box] == 0 {
					continue
				}

				for _, key := range keys[box] {
					status[key] = 1
				}
			}
			for i := 0; i < length; i++ {
				cur := boxs[0]
				boxs = boxs[1:]
				//if visited[cur] == 1 {
				//	continue
				//}

				if status[cur] == 1 {
					visited[cur] = 1
					total += candies[cur]
					boxs = append(boxs, containedBoxes[cur]...)
				}
			}
		}
	}
	bfs(initialBoxes)
	return total
}

func answerString1(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	res := ""
	n := len(word)
	maxLength := n - numFriends + 1
	maxIndex := [26][]int{}
	maxByte := uint8(0)
	for i := 0; i < n; i++ {
		b := word[i] - 'a'
		maxByte = max(maxByte, b)
		maxIndex[b] = append(maxIndex[b], i)
	}
	for _, i := range maxIndex[maxByte] {
		end := min(n, i+maxLength)
		res = max(res, word[i:end])
	}
	return res
}

// https://leetcode.cn/problems/find-the-lexicographically-largest-string-from-the-box-i/?envType=daily-question&envId=2025-06-04
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	res := ""
	n := len(word)
	maxLength := n - numFriends + 1
	for i := 0; i < n; i++ {
		res = max(res, word[i:min(i+maxLength, n)])
	}
	return res
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := [26]byte{}
	for i := range parent {
		parent[i] = byte(i)
	}
	var find func(byte) byte
	find = func(b byte) byte {
		if parent[b] != b {
			b = find(parent[b])
		}
		return
	}
	union := func(x, y byte) {
		small, big := find(x), find(y)
		if small > big {
			small, big = big, small
		}
		parent[big] = small
	}
	for i := 0; i < len(s1); i++ {
		union(s1[i]-'a', s2[i]-'a')
	}

	s := make([]byte, len(baseStr))
	for i, c := range baseStr {
		s[i] = find(byte(c)-'a') + 'a'
	}
	return string(s)

}
