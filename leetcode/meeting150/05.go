package meeting150

import (
	"strings"
)

/*
	哈希表
*/

func canConstruct(ransomNote string, magazine string) bool {
	need := [26]int{}
	cur := [26]int{}
	for _, word := range ransomNote {
		need[word-'a']++
	}
	for _, word := range magazine {
		cur[word-'a']++
	}
	for i := 0; i < 26; i++ {
		if need[i] > cur[i] {
			return false
		}
	}
	return true
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	maps := make(map[byte]byte)
	rmaps := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if v, ok := maps[s[i]]; ok {
			if v != t[i] {
				return false
			}
		} else {
			if _, ok2 := rmaps[t[i]]; ok2 {
				return false
			}
			maps[s[i]] = t[i]
			rmaps[t[i]] = s[i]
		}
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	maps := make(map[byte]string)
	rmaps := make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		if v, ok := maps[pattern[i]]; ok {
			if v != ss[i] {
				return false
			}
		} else {
			if _, ok2 := rmaps[ss[i]]; ok2 {
				return false
			}
			maps[pattern[i]] = ss[i]
			rmaps[ss[i]] = pattern[i]
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	need := [26]int{}
	cur := [26]int{}
	for _, word := range s {
		need[word-'a']++
	}
	for _, word := range t {
		cur[word-'a']++
	}
	for i := 0; i < 26; i++ {
		if need[i] > cur[i] {
			return false
		}
	}
	return true
}

func generateKey(num1 int, num2 int, num3 int) int {
	//res := []int{}
	base := 1
	ans := 0
	for num1 > 0 || num2 > 0 || num3 > 0 {
		ans += min(num1%10, num2%10, num3%10) * base
		num1 = num1 / 10
		num2 = num2 / 10
		num3 = num3 / 10
		base *= 10
	}
	//for len(res) != 4 {
	//	res = append(res, 0)
	//}
	//slices.Reverse(res)
	//ans := res[0]
	//for i := 1; i < 4; i++ {
	//	ans = ans*10 + res[i]
	//}
	return ans

}

func groupAnagrams(strs []string) [][]string {
	maps := make(map[[26]int][]string)
	for _, str := range strs {
		cur := [26]int{}
		for i := 0; i < len(str); i++ {
			cur[str[i]-'a']++
		}
		if _, ok := maps[cur]; ok {
			maps[cur] = append(maps[cur], str)
		} else {
			maps[cur] = []string{str}
		}
	}
	res := [][]string{}
	for _, v := range maps {
		res = append(res, v)
	}
	return res
}

// 两数之和
func isHappy(n int) bool {
	notHappy := make(map[int]bool)
	for {
		if notHappy[n] {
			return false
		}
		newn := getsum(n)
		if newn == 1 {
			return true
		} else {
			notHappy[n] = true
		}
		n = newn
	}
}

func getsum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func containsNearbyDuplicate(nums []int, k int) bool {
	index := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := index[nums[i]]; ok {
			if i-v <= k {
				return true
			}
		}
		index[nums[i]] = i
	}
	return false
}

// 最长连续序列
func longestConsecutive(nums []int) int {
	sets := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		sets[nums[i]] = true
	}
	res := 0
	for k, _ := range sets {
		if sets[k-1] {
			continue
		}
		cnt := 1
		cur := k
		for sets[cur+1] {
			cnt++
			cur++
		}
		res = max(cnt, res)
	}
	return res
}
