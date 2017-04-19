// https://leetcode.com/problems/longest-palindromic-substring/#/description
// thanks to @liji94188 :
// https://discuss.leetcode.com/topic/21848/ac-relatively-short-and-very-clear-java-solution
package leetcode

func longestPalindrome(s string) string {
	maxl := 0
	var result string
	l := len(s)
	if l < 2 {
		return s
	}
	for i := 0; i < l; i++ {
		if isPalindrome(s, i-maxl-1, i) {
			result = s[i-maxl-1 : i+1]
			maxl += 2
		} else if isPalindrome(s, i-maxl, i) {
			result = s[i-maxl : i+1]
			maxl += 1
		}
	}
	return result
}

func isPalindrome(s string, begin int, end int) bool {
	if begin < 0 {
		return false
	}
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}
