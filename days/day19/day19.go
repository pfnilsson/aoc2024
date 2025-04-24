package day19

import (
	"fmt"
	"log"
	"strings"

	"aoc2024/shared"
)

func parseTowelPatterns(towelPatterns []string) map[rune][]string {
	towelMap := make(map[rune][]string)

	for _, pattern := range towelPatterns {
		firstLetter := rune(pattern[0])
		towelMap[firstLetter] = append(towelMap[firstLetter], pattern)
	}
	return towelMap
}

func parseInput(inputRaw [][]string) (map[rune][]string, []string) {
	towelMap := parseTowelPatterns(strings.Split(inputRaw[0][0], ", "))
	return towelMap, inputRaw[1]
}

func checkDesign(design string, towelMap map[rune][]string, cache map[string]bool) bool {
	if val, ok := cache[design]; ok {
		return val
	}

	firstLetter := rune(design[0])
	if _, ok := towelMap[firstLetter]; !ok {
		cache[design] = false
		return false
	}

	for _, pattern := range towelMap[firstLetter] {
		if strings.HasPrefix(design, pattern) {
			if len(design) == len(pattern) {
				cache[design] = true
				return true
			}
			if checkDesign(design[len(pattern):], towelMap, cache) {
				cache[design] = true
				return true
			}
		}
	}
	cache[design] = false
	return false
}

func countValidCombinations(design string, towelMap map[rune][]string, cache map[string]int) int {
	if val, ok := cache[design]; ok {
		return val
	}

	if len(design) == 0 {
		return 0
	}

	firstLetter := rune(design[0])
	if _, ok := towelMap[firstLetter]; !ok {
		cache[design] = 0
		return 0
	}

	var combinations int
	for _, pattern := range towelMap[firstLetter] {
		if strings.HasPrefix(design, pattern) {
			if len(design) == len(pattern) {
				combinations++
			}
			combinations += countValidCombinations(design[len(pattern):], towelMap, cache)
		}
	}
	cache[design] = combinations
	return combinations
}

func part1(designs []string, towelMap map[rune][]string) {
	cache := make(map[string]bool)

	tot := 0
	for _, design := range designs {
		if checkDesign(design, towelMap, cache) {
			tot++
		}
	}
	fmt.Println("Part 1:", tot)
}

func part2(designs []string, towelMap map[rune][]string) {
	cache := make(map[string]int)

	tot := 0
	for _, design := range designs {
		tot += countValidCombinations(design, towelMap, cache)
	}
	fmt.Println("Part 2:", tot)
}

func Run() {
	lines, err := shared.ReadFileByBlankLine("days/day19/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	towelMap, designs := parseInput(lines)

	part1(designs, towelMap)
	part2(designs, towelMap)
}
