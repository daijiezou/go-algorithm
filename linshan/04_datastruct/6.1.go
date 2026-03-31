package _4_datastruct

import (
	"sort"
	"strings"
)

// 前缀树
type Trie struct {
	root *Node
}

type Node struct {
	son [26]*Node
	End bool
}

func Constructor2() Trie {
	return Trie{root: new(Node)}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, x := range word {
		node := cur.son[x-'a']
		if node == nil {
			node = new(Node)
			cur.son[x-'a'] = node
		}
		cur = cur.son[x-'a']
	}
	cur.End = true
}

func (this *Trie) find(word string) int {
	cur := this.root
	for _, x := range word {
		c := x - 'a'
		if cur.son[c] == nil {
			return 0
		}
		cur = cur.son[x-'a']
	}
	if cur.End {
		// 一模一样
		return 2
	}
	return 1

}

func (this *Trie) Search(word string) bool {
	return this.find(word) == 2
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

// 3597. 分割字符串
/*
给你一个字符串 s，按照以下步骤将其分割为 互不相同的段 ：

从下标 0 开始构建一个段。
逐字符扩展当前段，直到该段之前未曾出现过。
只要当前段是唯一的，就将其加入段列表，标记为已经出现过，并从下一个下标开始构建新的段。
重复上述步骤，直到处理完整个字符串 s。
返回字符串数组 segments，其中 segments[i] 表示创建的第 i 段。
*/
func partitionString(s string) []string {
	root := new(Node)
	cur := root
	res := []string{}
	//temp := make([]byte, 0)
	left := 0
	for i, x := range s {

		if cur.son[x-'a'] == nil {
			res = append(res, s[left:i+1])
			left = i + 1
			cur.son[x-'a'] = new(Node)
			cur = root
		} else {
			cur = cur.son[x-'a']
		}
	}
	return res
}

func partitionString2(s string) []string {
	preMap := make(map[string]bool)
	res := []string{}
	//temp := make([]byte, 0)
	left := 0
	for i := range s {
		t := s[left : i+1]
		if _, ok := preMap[t]; !ok {
			res = append(res, t)
			preMap[t] = true
			left = i + 1
		}
	}
	return res
}

// https://leetcode.cn/problems/replace-words/
// 648 单词替换
func replaceWords(dictionary []string, sentence string) string {
	dset := make(map[string]bool)
	for _, x := range dictionary {
		dset[x] = true
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if _, ok := dset[word[:j]]; ok {
				words[i] = word[:j]
				break
			}
		}
	}
	return strings.Join(words, " ")
}

// 720.词典中最长的单词
func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		a, b := words[i], words[j]
		return len(a) < len(b) || (len(a) == len(b) && a > b)
	})
	maps := make(map[string]struct{})
	res := ""
	maps[""] = struct{}{}
	for _, x := range words {
		if _, ok := maps[x[:len(x)-1]]; ok {
			res = x
			maps[x] = struct{}{}
		}

	}
	return res
}

/*
https://leetcode.cn/problems/sum-of-prefix-scores-of-strings/
*/

type TrieNode struct {
	son [26]*TrieNode
	Cnt int
}

func sumPrefixScores(words []string) []int {
	root := new(TrieNode)
	for _, word := range words {
		cur := root
		for _, x := range word {
			node := cur.son[x-'a']
			if node == nil {
				node = new(TrieNode)
				cur.son[x-'a'] = node
			}
			cur.son[x-'a'].Cnt++
			cur = cur.son[x-'a']
		}
	}

	ans := make([]int, len(words))
	for i, word := range words {
		n := len(word)
		score := 0
		cur := root
		for j := 0; j < n; j++ {
			cur = cur.son[word[j]-'a']
			score += cur.Cnt
		}
		ans[i] = score
	}
	return ans
}
