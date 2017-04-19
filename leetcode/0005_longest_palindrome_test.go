package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome_01(t *testing.T) {
	s := "babad"
	r := longestPalindrome(s)
	fmt.Println("longestPalindrome returned:", r)
	if r != "bab" && r != "aba" {
		t.FailNow()
	}
}

func TestLongestPalindrome_02(t *testing.T) {
	s := "cbbd"
	r := longestPalindrome(s)
	fmt.Println("longestPalindrome returned:", r)
	if r != "bb" {
		t.FailNow()
	}
}

func TestLongestPalindrome_03(t *testing.T) {
	s := "aaaaaaaaaaaaaaaaaabcaaaaaaaaaaaaaaaaa"
	r := longestPalindrome(s)
	fmt.Println("longestPalindrome returned:", r)
	if r != "aaaaaaaaaaaaaaaaaa" {
		t.FailNow()
	}
}
