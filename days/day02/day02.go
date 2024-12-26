package day02

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

func sameSign(a int, b int) bool {
	return (a > 0 && b > 0) || (a < 0 && b < 0)
}

func isValidSequence(a int, b int, refSign int) bool {
	diff := a - b

	if !sameSign(refSign, diff) {
		return false
	}

	absDiff := shared.AbsInt(diff)
	if absDiff < 1 || absDiff > 3 {
		return false
	}

	return true
}

func findRefSign(report []int) int {
	refSign0 := report[0] - report[1]
	refSign1 := report[1] - report[2]

	if sameSign(refSign0, refSign1) {
		return refSign0
	}

	refsign2 := report[2] - report[3]
	if sameSign(refSign0, refsign2) {
		return refSign0
	}

	return refSign1

}

func isSafe(report []int) bool {
	refSign := findRefSign(report)

	for i := 0; i < len(report)-1; i++ {
		if !isValidSequence(report[i], report[i+1], refSign) {
			return false
		}
	}
	return true
}

func isSafeWithDampener(report []int) bool {
	refSign := findRefSign(report)
	ignored := 0

	for i := 0; i < len(report)-1; i++ {
		if !isValidSequence(report[i], report[i+1], refSign) {
			ignored++
			if ignored > 1 {
				return false
			}

			// if the last element is invalid, we can ignore it
			if i == len(report)-2 {
				continue
			}

			// if the first element is invalid, we can ignore it
			if i == 0 && isValidSequence(report[i+1], report[i+2], refSign) {
				continue
			}

			// if the i+1 element is invalid, increment i to skip it
			if i+2 < len(report) && isValidSequence(report[i], report[i+2], refSign) {
				i++
			}

			// if the current element is invalid, proceed to ignore it
			if i > 0 && isValidSequence(report[i-1], report[i+1], refSign) {
				continue

			}
			return false
		}
	}
	return true
}

func part1(reports [][]int) {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

func part2(reports [][]int) {
	count := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}

func Run() {
	reports, err := shared.ReadFileByLineToSplitInts("days/day02/input.txt", " ")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(reports)
	part2(reports)
}
