package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumers01(t *testing.T) {
	a1 := &ListNode{3, nil}
	a2 := &ListNode{4, a1}
	a3 := &ListNode{2, a2}
	b1 := &ListNode{4, nil}
	b2 := &ListNode{6, b1}
	b3 := &ListNode{5, b2}
	result := addTwoNumbers(a3, b3)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Println()
}

func TestAddTwoNumers02(t *testing.T) {
	a1 := &ListNode{8, nil}
	b1 := &ListNode{1, nil}
	b2 := &ListNode{0, b1}
	result := addTwoNumbers(a1, b2)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Println()
}

func TestAddTwoNumers03(t *testing.T) {
	a1 := &ListNode{1, nil}
	b1 := &ListNode{9, nil}
	b2 := &ListNode{9, b1}
	b3 := &ListNode{9, b2}
	b4 := &ListNode{9, b3}
	result := addTwoNumbers(a1, b4)
	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
	fmt.Println()
}
