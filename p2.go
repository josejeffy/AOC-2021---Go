package main

import (
	"strings"
	// "fmt"
)

func p2_1(data string) int {
	depth := 0
	position := 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		instructions := strings.Split(strings.TrimSpace(line), " ")
		command := instructions[0]
		magnitude := ToInt(instructions[1])
		if command == "forward" {
			position = position + magnitude
		}
		if command == "down" {
			depth = depth + magnitude
		}
		if command == "up" {
			depth = depth - magnitude
		}
	}
	return depth * position
}

func p2_2(data string) int {
	depth := 0
	position := 0
	aim := 0
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		instructions := strings.Split(strings.TrimSpace(line), " ")
		command := instructions[0]
		magnitude := ToInt(instructions[1])
		if command == "forward" {
			position = position + magnitude
			depth = depth + aim*magnitude
		}
		if command == "down" {
			aim = aim + magnitude
		}
		if command == "up" {
			aim = aim - magnitude
		}
	}
	return depth * position
}
