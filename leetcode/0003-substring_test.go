package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring01(t *testing.T) {
	s := "abcabcbb"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 3 {
		t.FailNow()
	}
	fmt.Println("substring length =", length)
}
func TestLengthOfLongestSubstring02(t *testing.T) {
	s := "pwwkew"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 3 {
		t.FailNow()
	}
	fmt.Println("substring length =", length)
}
func TestLengthOfLongestSubstring03(t *testing.T) {
	s := "c"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 1 {
		t.FailNow()
	}
	fmt.Println("substring length =", length)
}
func TestLengthOfLongestSubstring04(t *testing.T) {
	s := "au"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 2 {
		t.FailNow()
	}
	fmt.Println("substring length =", length)
}
func TestLengthOfLongestSubstring05(t *testing.T) {
	s := "dvdf"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 3 {
		t.FailNow()
	}
	fmt.Println("substring length = ", length)
}
func TestLengthOfLongestSubstring06(t *testing.T) {
	s := "abba"
	fmt.Println("orig str is", s)
	length := lengthOfLongestSubstring(s)
	if length != 2 {
		t.FailNow()
	}
	fmt.Println("substring length =", length)
}
