package sort

func BubbleSort(array []int64) {
	length := len(array)
	for i := 1; i < length; i++ {
		for j := 0; j < length-i; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
}
