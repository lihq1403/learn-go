package main

import "fmt"

/**
寻找最长不含有重复字符的字串
*/
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("哈哈哈哈哈你好"))
	//fmt.Println(lengthOfNonRepeatingSubStr("bbbbb"))
	//fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	//fmt.Println(lengthOfNonRepeatingSubStr(""))
	//fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	//fmt.Println(lengthOfNonRepeatingSubStr("a"))
}
