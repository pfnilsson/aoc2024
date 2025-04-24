package day01

import (
	"fmt"
	"log"
	"sort"

	"aoc2024/shared"
)

func parseLists(rawInput [][]int) ([]int, []int, error) {
	var list1, list2 []int

	for _, line := range rawInput {
		list1 = append(list1, line[0])
		list2 = append(list2, line[1])
	}

	return list1, list2, nil
}

func sortLists(list1 []int, list2 []int) {
	sort.Ints(list1)
	sort.Ints(list2)
}

func part1(list1 []int, list2 []int) {
	sortLists(list1, list2)

	totalDistances := 0
	for i := range list1 {
		totalDistances += shared.AbsInt(list1[i] - list2[i])
	}

	fmt.Println("Part 1:", totalDistances)
}

func part2(list1 []int, list2 []int) {
	counts := make(map[int]int)
	for _, val := range list2 {
		counts[val]++
	}

	simScore := 0
	for _, val := range list1 {
		simScore += val * counts[val]
	}

	fmt.Println("Part 2:", simScore)
}

func Run() {
	rawInput, err := shared.ReadFileByLineToSplitInts("days/day01/input.txt", "   ")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	list1, list2, err := parseLists(rawInput)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(list1, list2)
	part2(list1, list2)
}
