package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(input)
	score := 0
	for _, val := range strings.Split(string(input), "\n") {
		left := getVal(strings.Split(val, " ")[0])
		right := getVal(strings.Split(val, " ")[1])
		if right == 1 {
			fmt.Println("loose")

			if (left - 1) != 0 {
				score += left - 1
			} else {
				score += 3
			}

			//fmt.Println(left-1) )
		}
		if right == 2 {
			fmt.Println("draw")
			score += 3
			score += left
		}
		if right == 3 {
			fmt.Println("win")
			score += 6
			if (left + 1) != 4 {
				score += left + 1
			} else {
				score += 1
			}
		}
		fmt.Println(score)

	}
	fmt.Println(score)
}
func calcPart1() {
	score := 0
	for _, val := range strings.Split(string(input), "\n") {
		left := getVal(strings.Split(val, " ")[0])
		right := getVal(strings.Split(val, " ")[1])

		fmt.Println(left, right)
		if left-right == 0 {
			fmt.Println("draw")
			score += 3
			score += right
		}
		if left-right == -1 || left-right == 2 {
			fmt.Println("win")
			score += 6
			score += left
		}
		if left-right == 1 || left-right == -2 {
			fmt.Println("loose")
			score += right
		}
	}
	fmt.Println(score)

}

func getVal(s string) int {
	if s == "A" || s == "X" {
		return 1
	}
	if s == "B" || s == "Y" {
		return 2
	}
	if s == "C" || s == "Z" {
		return 3
	}
	panic("err")
}

/*
	A = ROCK	1
	B = PAPER	2
	C SCISSOR	3

	X = ROCK	1
	Y = PAPER	2
	Z = SCISSOR 3

	left - right = 0 = draw
	left - right = -1 = win
	left - right = -2 = loose

	left - right = 1 = loose
	left - right = 2 = win
*/
