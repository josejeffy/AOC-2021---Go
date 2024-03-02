package main

import (
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.Atoi(strings.TrimSpace(s))
	return i
}
