package main

import "fmt"

//给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//进阶：
//如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
//致谢：
//特别感谢 @pbrother 添加此问题并且创建所有测试用例。
//示例 1：
//输入：s = "abc", t = "ahbgdc"
//输出：true
//示例 2：
//输入：s = "axc", t = "ahbgdc"
//输出：false
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/is-subsequence
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

//思路: 把t打成chars 和 s的chars 一一比较
// 要保证顺序的话 要记录 index
// 后面字母的index必须必前面的index大
func isSubsequence(s string, t string) bool {
	var rst bool
	var idx int
	for k, tc := range t {
	continueTag:
		for _, sc := range s {
			if sc == tc { //有相同的字符
				if k < idx { //
					rst = false
					return rst
				} else {
					idx = k
					rst = true
					continue continueTag
				}
			}
			fmt.Println(sc, tc)
		}
	}

	return false
}

func main() {
	r := isSubsequence("abc", "ahbgdc")
	fmt.Println(r)
}
