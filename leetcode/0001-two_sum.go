package leetcode

func TwoSum(nums []int, target int) []int {
	var first int
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		first = nums[i]
		for j := i + 1; j < numsLen; j++ {
			if first+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
