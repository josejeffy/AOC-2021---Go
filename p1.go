package main

import (
	"strings"
)


func p1(data string, windowWidth int) int {
	lines := strings.Split(data, "\n")
	// windowWidth := 3
	left := 0
	right := left + windowWidth - 1
	prevSum := 0
	count := 0
	
	for right < len(lines) {
		windowSum := 0
		for i := left; i <= right; i++ {
			windowSum = windowSum + ToInt(lines[i])
		}
		if left != 0 && windowSum > prevSum {
			count = count + 1
		}
		prevSum = windowSum
		right = right + 1
		left = left + 1
	}
	return count
}

func p1_1(data string) int {
	return p1(data,1)
}

func p1_2(data string) int {
	return p1(data,3)
}
