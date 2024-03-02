package main

import (
	// "fmt"
	"math"
	"strings"
)

func p3_1(data string) int {
	counter := make(map[int]int)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		for i := 0; i < len(strings.TrimSpace(line)); i++ {
			if string(line[i]) == "1" {
				counter[i] = counter[i] + 1
			}
		}
	}
	gamma := 0
	epsilon := 0
	n_digits := len(counter)
	threshold := len(lines)/2
	for c := range n_digits {
		inc := int(math.Pow(2, float64(n_digits-c-1)))
		if counter[c] > threshold{
			gamma = gamma + inc
		} else {
			epsilon = epsilon + inc
		}
	}
	return gamma * epsilon
}
