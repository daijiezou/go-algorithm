package _4_datastruct

// 单调栈

// 739. 每日温度
/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，
下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
*/
func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	for i, x := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < x {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[pop] = i - pop
		}
		stack = append(stack, i)
	}
	return res
}

func dailyTemperaturesReverse(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] <= temperatures[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 { //后面没有比我更大的

		} else {
			nextG := stack[len(stack)-1]
			res[i] = nextG - i
		}
		stack = append(stack, i)
	}
	return res
}
