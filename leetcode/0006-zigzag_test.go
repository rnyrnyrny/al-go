package leetcode

import (
	"fmt"
	"testing"
)

func TestZigZag_01(t *testing.T) {
	s := "ahlqcxrzmqebfvgpvjcvpupqfvnlbiodsatuznvldcoixuxudcpvwelbcbodjejdecxgpttuviduecokeefaks"
	n := 48
	expected := "ahlqcxrzmqsekbaffvegepkvojccevupduipvquftvtnplgbxicoeddsjaetjudzonbvclbdlceowivxpucxdu"
	result := Convert(s, n)
	if result != expected {
		fmt.Println("zigzag conversion result error..")
		fmt.Println("error result is ", result)
		t.FailNow()
	}
}

func TestZigZag_02(t *testing.T) {
	s := "A"
	n := 1
	expected := "A"
	result := Convert(s, n)
	if result != expected {
		fmt.Println("zigzag conversion result error..")
		fmt.Println("error result is ", result)
		t.FailNow()
	}
}

func TestZigZag_03(t *testing.T) {
	s := "ABCDE"
	n := 2
	expected := "ACEBD"
	result := Convert(s, n)
	if result != expected {
		fmt.Println("zigzag conversion result error..")
		fmt.Println("error result is ", result)
		t.FailNow()
	}
}
