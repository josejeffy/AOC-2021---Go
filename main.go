package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("p1.txt")
	input := string(data)
	fmt.Println(p1_1(input))
	fmt.Println(p1_2(input))
}
