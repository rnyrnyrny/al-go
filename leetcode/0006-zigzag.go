// zig-zag conversion
// https://leetcode.com/problems/zigzag-conversion/#/description
// Thanks to tju_xu97
// https://discuss.leetcode.com/topic/21196/a-10-lines-one-pass-o-n-time-o-1-space-accepted-solution-with-detailed-explantation/4

package leetcode

func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	var result string
	// cycle means a zig and a zag
	cycle := 2*numRows - 2
	// iterate by row
	for i := 0; i < numRows; i++ {
		// for each row iterate by cycle
		for j := i; j < len(s); j += cycle {
			result += string(s[j])
			j2 := j - i + cycle - i
			if i != 0 && i != numRows-1 && j2 < len(s) {
				result += string(s[j2])
			}
		}
	}
	return result
}
