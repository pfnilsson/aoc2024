package day05

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type tuple struct {
	a string
	b string
}

type sorter struct {
	values []string
	rules  *shared.Set[tuple]
}

func (s sorter) Len() int {
	return len(s.values)
}

func (s sorter) Less(i, j int) bool {
	return s.rules.Contains(tuple{a: s.values[i], b: s.values[j]})
}

func (s sorter) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func parseInput(rawInput []string) (rules *shared.Set[tuple], pageCollection [][]string) {
	rules = shared.NewSet[tuple]()

	blankFound := false
	for _, line := range rawInput {
		if line == "" {
			blankFound = true
			continue
		}

		if !blankFound {
			splitLine := strings.Split(line, "|")
			rules.Add(tuple{a: splitLine[0], b: splitLine[1]})
		} else {
			pagesStr := strings.Split(line, ",")

			var pages []string
			for _, page := range pagesStr {
				pages = append(pages, page)
			}
			pageCollection = append(pageCollection, pages)
		}
	}
	return rules, pageCollection
}

func calculateMiddleIndexTotal(pages [][]string) (int, error) {
	tot := 0
	for _, page := range pages {
		middleIndex := len(page) / 2
		middlevalue, err := strconv.Atoi(page[middleIndex])
		if err != nil {
			return 0, err
		}
		tot += middlevalue
	}
	return tot, nil
}

func part(pages [][]string, partNr int) error {
	tot, err := calculateMiddleIndexTotal(pages)
	if err != nil {
		return err
	}
	fmt.Printf("Part %v: %v\n", partNr, tot)
	return nil
}

func part1(pages [][]string) error {
	return part(pages, 1)
}

func part2(pages [][]string) error {
	return part(pages, 2)
}

func sortPages(rules *shared.Set[tuple], pageCollection [][]string) (correct [][]string, incorrect [][]string) {
	s := sorter{rules: rules}

	for _, pageNumbers := range pageCollection {
		pageNumbersSorted := make([]string, len(pageNumbers))

		copy(pageNumbersSorted, pageNumbers)
		s.values = pageNumbersSorted
		sort.Sort(s)

		if shared.SlicesEqual(pageNumbers, pageNumbersSorted) {
			correct = append(correct, pageNumbersSorted)
		} else {
			incorrect = append(incorrect, pageNumbersSorted)
		}
	}
	return correct, incorrect
}

func Run() {
	rawInput, err := shared.ReadFileByLine("days/day05/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	rules, pageCollection := parseInput(rawInput)
	correct, incorrect := sortPages(rules, pageCollection)

	err = part1(correct)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = part2(incorrect)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
