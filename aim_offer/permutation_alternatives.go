package aim_offer

import "sort"

// 方法2：基于字符计数的回溯（避免重复）
func PermutationByCount(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	// 统计每个字符的出现次数
	charCount := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		charCount[str[i]]++
	}

	// 获取所有不重复的字符并排序
	var chars []byte
	for char := range charCount {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	var result []string
	var track []byte

	var backtrack func()
	backtrack = func() {
		if len(track) == len(str) {
			result = append(result, string(track))
			return
		}

		for _, char := range chars {
			if charCount[char] > 0 {
				// 使用这个字符
				charCount[char]--
				track = append(track, char)
				backtrack()
				// 回溯
				track = track[:len(track)-1]
				charCount[char]++
			}
		}
	}

	backtrack()
	return result
}

// 方法3：使用next_permutation算法（字典序）
func PermutationLexicographic(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	chars := []byte(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	var result []string
	result = append(result, string(chars))

	for nextPermutation(chars) {
		result = append(result, string(chars))
	}

	return result
}

// 辅助函数：生成下一个字典序排列
func nextPermutation(nums []byte) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}

	// 1. 从右往左找第一个升序对 (i, i+1)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i < 0 {
		return false // 已经是最大排列
	}

	// 2. 从右往左找第一个大于nums[i]的数
	j := n - 1
	for nums[j] <= nums[i] {
		j--
	}

	// 3. 交换nums[i]和nums[j]
	nums[i], nums[j] = nums[j], nums[i]

	// 4. 反转i+1到末尾的部分
	reverse2(nums[i+1:])

	return true
}

// 辅助函数：反转数组
func reverse2(arr []byte) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// 方法4：使用位运算（仅适用于小规模且无重复字符）
func PermutationBitMask(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	// 检查是否有重复字符
	charSet := make(map[byte]bool)
	for i := 0; i < len(str); i++ {
		if charSet[str[i]] {
			// 有重复字符，回退到方法1
			return Permutation(str)
		}
		charSet[str[i]] = true
	}

	n := len(str)
	var result []string

	var backtrack func(track []byte, used int)
	backtrack = func(track []byte, used int) {
		if len(track) == n {
			result = append(result, string(track))
			return
		}

		for i := 0; i < n; i++ {
			// 检查第i位是否已使用
			if (used & (1 << i)) != 0 {
				continue
			}

			track = append(track, str[i])
			backtrack(track, used|(1<<i))
			track = track[:len(track)-1]
		}
	}

	backtrack([]byte{}, 0)
	return result
}

// 性能比较函数
func ComparePermutationMethods(str string) {
	println("=== 全排列方法性能比较 ===")
	println("输入字符串:", str)

	// 方法1：排序+跳过重复
	result1 := Permutation(str)
	println("方法1 (排序+跳过重复):", len(result1), "个结果")

	// 方法2：字符计数
	result2 := PermutationByCount(str)
	println("方法2 (字符计数):", len(result2), "个结果")

	// 方法3：字典序
	result3 := PermutationLexicographic(str)
	println("方法3 (字典序):", len(result3), "个结果")

	// 验证结果是否一致
	if len(result1) == len(result2) && len(result2) == len(result3) {
		println("✅ 所有方法结果数量一致")
	} else {
		println("❌ 方法结果不一致")
	}
}
