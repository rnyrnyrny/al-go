package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseInteger_01(t *testing.T) {
	result := reverse(1)
	if result != 1 {
		fmt.Println("reverse returned", result)
		t.FailNow()
	}
}

func TestReverseInteger_02(t *testing.T) {
	result := reverse(123)
	if result != 321 {
		fmt.Println("reverse returned", result)
		t.FailNow()
	}
}

func TestReverseInteger_03(t *testing.T) {
	result := reverse(10)
	if result != 1 {
		fmt.Println("reverse returned", result)
		t.FailNow()
	}
}

func TestReverseInteger_04(t *testing.T) {
	result := reverse(-1234)
	if result != -4321 {
		fmt.Println("reverse returned", result)
		t.FailNow()
	}
}

func TestReverseInteger_05(t *testing.T) {
	result := reverse(1534236469)
	if result != 0 {
		fmt.Println("reverse returned", result)
		t.FailNow()
	}
}
