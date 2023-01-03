package main

import (
	"fmt"
	"github.com/slx618/day09/03unit_test/split_string"
)

// 以 _test.go 结尾
// Test 逻辑行为是否正确
// Benchmark 性能测试
// Example 文档示例
func main() {
	c := split_string.Split("abcabcab", "a")
	fmt.Println(c)
}
