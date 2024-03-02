package main

import (
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic("Error")
	}
}

func ToInt(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	check(err)
	return i
}
