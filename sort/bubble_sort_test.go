package sort

import (
	"fmt"
	"testing"

	"github.com/rnyrnyrny/al-go/utils"
)

func TestBubbleSort(t *testing.T) {
	arr := utils.RandomArray(10)
	BubbleSort(arr)
	fmt.Println(arr)
	length := len(arr)
	for i := 0; i < length-1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("BubbleSort failed..")
			t.FailNow()
		}
	}
}
