// https://leetcode.com/problems/add-two-numbers/#/description

package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// This code is ungly..
func addTwoNumbersUgly(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1a, l2a, longer, re []int
	var pos, result *ListNode
	var carry int = 0
	var len1, len2, base, total int
	pos = l1
	for pos.Next != nil {
		l1a = append(l1a, pos.Val)
		pos = pos.Next
	}
	l1a = append(l1a, pos.Val)
	pos = l2
	for pos.Next != nil {
		l2a = append(l2a, pos.Val)
		pos = pos.Next
	}
	l2a = append(l2a, pos.Val)
	fmt.Println(l1a)
	fmt.Println(l2a)
	len1 = len(l1a)
	len2 = len(l2a)
	if len1 != len2 {
		if len1 > len2 {
			base = len2
			total = len1
			longer = l1a
		} else {
			base = len1
			total = len2
			longer = l2a
		}
		re = make([]int, total)
	} else {
		base = len1
		total = len1 + 1
		re = make([]int, total)
	}
	fmt.Println("total=", total)
	for i := 0; i < base; i++ {
		tmp := l1a[i] + l2a[i] + carry
		fmt.Println("tmp=", tmp)
		if tmp > 9 {
			re[total-i-1] = tmp % 10
			carry = 1
		} else {
			re[total-i-1] = tmp
			carry = 0
		}
	}
	if len1 == len2 {
		re[total-1-base] = carry
		carry = 0
	} else {
		for i := base; i < total; i++ {
			tmp := longer[i] + carry
			if tmp > 9 {
				re[total-i-1] = tmp % 10
				carry = 1
			} else {
				re[total-i-1] = tmp
				carry = 0
			}
		}
	}
	if carry > 0 {
		re = append([]int{carry}, re...)
	}
	fmt.Println(re)
	if re[0] == 0 {
		re = re[1:]
	}
	fmt.Println(re)
	result = &ListNode{re[len(re)-1], nil}
	pos = result
	for i := 1; i < len(re); i++ {
		result.Next = &ListNode{re[len(re)-i-1], nil}
		result = result.Next
	}
	fmt.Println(pos)
	return pos
}

// This solution is by @potpie
// https://discuss.leetcode.com/topic/799/is-this-algorithm-optimal-or-what
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum int = 0
	l1c := l1
	l2c := l2
	result := &ListNode{0, nil}
	tmp := result
	for l1c != nil || l2c != nil {
		if l1c != nil {
			sum += l1c.Val
			l1c = l1c.Next
		}
		if l2c != nil {
			sum += l2c.Val
			l2c = l2c.Next
		}
		// this is faster than sum%10
		tmp.Next = &ListNode{sum - sum/10*10, nil}
		tmp = tmp.Next
		sum /= 10
	}
	if sum == 1 {
		tmp.Next = &ListNode{1, nil}
	}
	return result.Next
}
