package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8
	result := TwoSum(nums, target)
	if result == nil {
		fmt.Println("TwoSum failed: cannot find answer")
		t.FailNow()
	}
	if result[0] != 2 || result[1] != 4 {
		fmt.Println("TwoSum failed: wrong answer")
		t.FailNow()
	}
}
