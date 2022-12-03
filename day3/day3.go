package main

import (
	_ "embed"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"golang.org/x/exp/slices"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(input, "\n")

	var alphabet []rune
	for r := 'a'; r <= 'z'; r++ {
		alphabet = append(alphabet, r)
	}
	for r := 'A'; r <= 'Z'; r++ {
		alphabet = append(alphabet, r)
	}
	allRows := strings.Split(string(input), "\n")
	totalScore := 0
	for i, _ := range allRows {
		if (i+1)%3 == 0 {
			elf1 := mapset.NewSetFromSlice(toInterface(strings.Split(allRows[i-2], "")))
			elf2 := mapset.NewSetFromSlice(toInterface(strings.Split(allRows[i-1], "")))
			elf3 := mapset.NewSetFromSlice(toInterface(strings.Split(allRows[i], "")))
			common := elf1.Intersect(elf2).Intersect(elf3)
			fmt.Println(common)
			char := fmt.Sprintf("%v", common.Pop())
			score := slices.Index(alphabet, []rune(char)[0])
			fmt.Println(score)
			totalScore += score + 1
		}
	}
	fmt.Println(totalScore)
}

func toInterface(runes []string) []interface{} {
	interfaceList := make([]interface{}, len(runes))
	for i, v := range runes {
		interfaceList[i] = v
	}
	return interfaceList
}

func part1(alphabet []rune) {

	fmt.Println(string(alphabet))
	score := 0
	for _, val := range strings.Split(string(input), "\n") {
		chars := []rune(val)
		left := chars[0 : len(chars)/2]
		right := chars[len(chars)/2 : len(chars)]
		fmt.Println(string(left))
		fmt.Println(string(right))
		set := make(map[rune]bool)
		for i := range left {
			if slices.Contains(left, right[i]) {
				set[right[i]] = true
			}
		}
		fmt.Println("present in both: {}", set)
		var priority int
		for s := range set {
			priority = slices.Index(alphabet, s)
			fmt.Println("prio", priority)
			score += priority + 1
		}

	}
	fmt.Println("score: ", score)
}
