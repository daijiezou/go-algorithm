package _7game

// https://leetcode.cn/problems/burst-balloons/description/
/*
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。

现在要求你戳破所有的气球。戳破第 i 个气球，你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。 这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。

求所能获得硬币的最大数量。
*/
var maxScore int

func maxCoins(nums []int) int {
	maxCoinsBackTrack(nums, 0)
	return maxScore
}

func maxCoinsBackTrack(nums []int, score int) int {
	/*
	 if (nums 为空) {
	        res = max(res, score);
	        return;
	    }
	    for (int i = 0; i < nums.length; i++) {
	        int point = nums[i-1] * nums[i] * nums[i+1];
	        int temp = nums[i];
	        // 做选择
	        在 nums 中删除元素 nums[i]
	        // 递归回溯
	        backtrack(nums, score + point);
	        // 撤销选择
	        将 temp 还原到 nums[i]
	    }
	*/
	return 0
}
