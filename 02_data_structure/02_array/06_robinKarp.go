package _2_array

import "math"

// https://leetcode.cn/problems/repeated-dna-sequences/submissions/533765602/
func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	nums := make([]int, len(s))
	for i := 0; i < len(nums); i++ {
		switch s[i] {
		case 'A':
			nums[i] = 0
		case 'G':
			nums[i] = 1
		case 'C':
			nums[i] = 2
		case 'T':
			nums[i] = 3
		}
	}
	left, right := 0, 0
	seen := make(map[int]bool)
	resMap := make(map[string]struct{})
	res := make([]string, 0)

	// 数字位数
	L := 10
	// 进制
	R := 4
	// 存储 R^(L - 1) 的结果
	RL := int(math.Pow(float64(R), float64(L-1)))
	// 维护滑动窗口中字符串的哈希值
	windowHash := 0

	for right < length {
		windowHash = R*windowHash + nums[right]
		right++
		if right-left == 10 {
			if _, ok := seen[windowHash]; ok {
				resMap[s[left:right]] = struct{}{}
			} else {
				seen[windowHash] = true
			}
			windowHash = windowHash - nums[left]*RL
			left++
		}
	}
	for k := range resMap {
		res = append(res, k)
	}
	return res
}

func transStr(in []byte) string {
	var res string
	for i := 0; i < len(in); i++ {
		res += string(in[i])
	}
	return res
}

// Rabin-Karp指纹字符串查找算法
func RabinKarp(txt string, pat string) int {
	// 数字位数
	L := len(pat)
	// 进制 (只考虑 ASCII 编码)
	R := int64(256)
	// 取一个比较大的素数作为求模的除数
	Q := int64(1658598167)
	// R^(L - 1) 的结果
	var RL int64 = 1
	for i := 1; i <= L-1; i++ {
		// 计算过程中不断求模，避免溢出
		RL = (RL * R) % Q
	}

	var patHash int64
	for i := 0; i < len(pat); i++ {
		patHash = (R*patHash + int64(pat[i])) % Q
	}
	// 维护滑动窗口中字符串的哈希值
	windowHash := int64(0)
	left, right := 0, 0
	length := len(txt)
	for right < length {
		windowHash = (R*windowHash)%Q + int64(txt[right])%Q
		right++
		if right-left == L {
			if windowHash == patHash {
				if pat == txt[left:right] {
					return left
				}
			}
			// 缩小窗口
			windowHash = ((windowHash - int64(txt[left])*RL%Q) + Q) % Q
			// 因为 windowHash - (txt[left] * RL) % Q 可能是负数
			// 所以额外再加一个 Q，保证 windowHash 不会是负数
		}
	}
	return -1
}
