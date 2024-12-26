package day25

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

func parseInput(inputRaw [][]string) ([][5]int, [][5]int) {
	var keys [][5]int
	var locks [][5]int
	var isKey bool

	for _, itemRaw := range inputRaw {
		item := [5]int{-1, -1, -1, -1, -1}
		if itemRaw[0][0] == '#' {
			isKey = true
		} else {
			isKey = false
		}

		for _, row := range itemRaw {
			for i, char := range row {
				if char == '#' {
					item[i]++
				}
			}
		}

		if isKey {
			keys = append(keys, item)
		} else {
			locks = append(locks, item)
		}
	}

	return keys, locks
}

func validCombo(key [5]int, lock [5]int) bool {
	for i := 0; i < 5; i++ {
		if key[i]+lock[i] > 5 {
			return false
		}
	}
	return true
}

func part1(keys [][5]int, locks [][5]int) {
	tot := 0
	for _, key := range keys {
		for _, lock := range locks {
			if validCombo(key, lock) {
				tot++
			}
		}
	}
	fmt.Println("Part 1:", tot)
}

func Run() {
	inputRaw, err := shared.ReadFileByBlankLine("days/day25/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	keys, locks := parseInput(inputRaw)

	part1(keys, locks)
}
