// https://leetcode.com/problems/longest-substring-without-repeating-characters/#/description
// thanks to @tongxing
// https://discuss.leetcode.com/topic/20192/4ms-c-code-in-12-lines
package leetcode

func lengthOfLongestSubstring(s string) int {
	var tmp int
	m := make([]int, 128)
	length := 0
	start := 0
	// end is the sentinal
	end := 0
	for end < len(s) {
		tmp = m[s[end]]
		// use 1 to len(s) to mark every char
		m[s[end]] = end + 1
		// if tmp is larger than 0, that
		// means this char has shown before
		// specifically, tmp should be larger
		// than start
		if tmp > start {
			if (end - start) > length {
				length = end - start
			}
			// tmp is end+1 so no need to
			// +1
			start = tmp
		}
		end++
	}
	if (end - start) > length {
		length = end - start
	}
	return length
}
