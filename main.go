package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("p2.txt")
	input := string(data)
	fmt.Println(p2_1(input))
	fmt.Println(p2_2(input))
}
