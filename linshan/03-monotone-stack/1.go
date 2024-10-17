package _3_monotone_stack

// 从左到右遍历
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := 0; i < length; i++ {
		for len(s) > 0 && temperatures[i] > temperatures[s[len(s)-1]] {
			j := s[len(s)-1]
			s = s[:len(s)-1]
			res[j] = i - j
		}
		s = append(s, i)
	}
	return res
}

// 从右向左
func dailyTemperatures2(temperatures []int) []int {
	length := len(temperatures)
	res := make([]int, length)
	s := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i] = 0
		} else {
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}
