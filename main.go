package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("p3.txt")
	input := string(data)
	fmt.Println(p3_1(input))
	// fmt.Println(p3_2(input))
}
