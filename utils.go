package utils

import "math/rand"

func RandomArray(length uint32) []int64 {
	arr := make([]int64, length)
	var i uint32
	for i = 0; i < length; i++ {
		arr[i] = rand.Int63()
	}
	return arr
}
