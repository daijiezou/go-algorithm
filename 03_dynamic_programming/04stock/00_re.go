package _4stock

// 正则表达式
func isMatch(s string, p string) bool {
	// 备忘录
	memo := make([][]int, len(s))
	for i := range memo {
		memo[i] = make([]int, len(p))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 指针 i，j 从索引 0 开始移动
	return dp(s, 0, p, 0, memo)
}

/*
计算 p[j..] 是否匹配 s[i..]
.a*b" 就可以匹配文本 "zaaab"，、也可以匹配 "cb"；
模式串 "a..b" 可以匹配文本 "amnb"；
模式串 ".*" 可以匹配任何文本。
*/

func dp(strs string, i int, strp string, j int, memo [][]int) bool {
	if j == len(strp) {
		return i == len(strs)
	}

	if i == len(strs) {
		// 查看后续strp能否匹配空串
		// 一定是a*b*c*这种才能匹配空串
		if (len(strp)-j)%2 == 1 {
			return false
		}

		for ; j < len(strp)-1; j += 2 {
			if strp[j+1] != '*' {
				return false
			}
		}
		return true
	}

	if memo[i][j] != -1 {
		// 查备忘录，防止重复计算
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}
	}
	res := false
	// 匹配成功，或者strp[j]
	if strs[i] == strp[j] || strp[j] == '.' {
		if j < len(strp)-1 && strp[j+1] == '*' {
			// 匹配0次或者无数次
			res = dp(strs, i, strp, j+2, memo) || dp(strs, i+1, strp, j, memo)
		} else {
			// 正常匹配一次
			res = dp(strs, i+1, strp, j+1, memo)
		}
	} else {
		if j < len(strp)-1 && strp[j+1] == '*' {
			//匹配0次
			res = dp(strs, i, strp, j+2, memo)
		} else {
			res = false
		}
	}
	if res {
		memo[i][j] = 1
	}
	return res
}
