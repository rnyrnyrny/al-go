// https://leetcode.com/problems/reverse-integer/#/description
package leetcode

import "math"

// terrible code
func terribleReverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 || x == 0 {
		return 0
	}
	tmp := 0
	result := 0
	divider := 10
	digits := []int{0}
	if x > 0 {
		tmp = x
	} else {
		tmp = -x
	}
	if tmp < divider {
		return x
	}
	for tmp >= divider {
		digits = append(digits, tmp-tmp/divider*divider)
		//fmt.Println("digits:", digits)
		tmp = tmp / divider
	}
	digits = append(digits, tmp)
	//fmt.Println("digits:", digits)
	divider = 1
	for i := 0; i < len(digits)-1; i++ {
		result += digits[len(digits)-1-i] * divider
		divider *= 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if x < 0 {
		result = -result
	}
	return result
}

func reverse(x int) int {
	var result int32 = 0
	var newResult int32 = 0
	var tmp int32 = int32(x)
	for tmp != 0 {
		mod := tmp - tmp/10*10
		newResult = result*10 + mod
		if (newResult-mod)/10 != result {
			return 0
		}
		//if newResult > math.MaxInt32 || newResult < math.MinInt32 {
		//return 0
		//}
		result = newResult
		tmp = tmp / 10
	}
	return int(result)
}
