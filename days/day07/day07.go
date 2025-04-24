package day07

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2024/shared"
)

type testCase struct {
	testValue  int
	components []int
}

func newTestCase(testValue int, components []int) testCase {
	return testCase{
		testValue:  testValue,
		components: components,
	}
}

func parseLine(line string) (testCase, error) {
	splitLine := strings.Split(line, ":")
	testValueStr := splitLine[0]
	testValue, err := strconv.Atoi(testValueStr)
	if err != nil {
		return testCase{}, err
	}

	componentsStr := strings.Split(splitLine[1], " ")
	var components []int
	for _, componentStr := range componentsStr {
		if componentStr == "" {
			continue
		}

		component, err := strconv.Atoi(componentStr)
		if err != nil {
			return testCase{}, err
		}
		components = append(components, component)
	}
	return newTestCase(testValue, components), nil
}

func reverseConcat(a, b int) (int, bool) {
	numDigits := 0
	tempB := b
	for tempB > 0 {
		numDigits++
		tempB /= 10
	}

	aLastDigits := a % pow(10, numDigits)
	if aLastDigits != b {
		return 0, false
	}

	return a / pow(10, numDigits), true
}

func pow(base int, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func evaluateTestCase(tc testCase, conc bool) bool {
	reducedValue := tc.testValue
	for i := len(tc.components) - 1; i >= 0; i-- {
		component := tc.components[i]

		if reducedValue%component == 0 && i > 0 {
			dividedValue := reducedValue / component
			childCase := newTestCase(dividedValue, tc.components[:i])
			if evaluateTestCase(childCase, conc) {
				return true
			}
		}

		if conc {
			if unConcatenated, ok := reverseConcat(reducedValue, component); ok {
				childCase := newTestCase(unConcatenated, tc.components[:i])
				if evaluateTestCase(childCase, conc) {
					return true
				}
			}
		}

		reducedValue = reducedValue - component

		if reducedValue < 0 {
			return false
		}
	}
	return reducedValue == 0
}

func parseTestCases(lines []string) ([]testCase, error) {
	var testCases []testCase
	for _, line := range lines {
		tc, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		testCases = append(testCases, tc)
	}
	return testCases, nil
}

func part1(testCases []testCase) (int, []testCase) {
	var (
		failedCases []testCase
		tot         int
	)

	for _, tc := range testCases {
		if evaluateTestCase(tc, false) {
			tot += tc.testValue
		} else {
			failedCases = append(failedCases, tc)
		}
	}

	fmt.Println("Part 1:", tot)
	return tot, failedCases
}

func part2(subtotal int, testCases []testCase) {
	for _, tc := range testCases {
		if evaluateTestCase(tc, true) {
			subtotal += tc.testValue
		}
	}
	fmt.Println("Part 2:", subtotal)
}

func Run() {
	lines, err := shared.ReadFileByLine("days/day07/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	testCases, err := parseTestCases(lines)
	if err != nil {
		log.Fatalf("Error parsing test cases: %v", err)
		return
	}

	subTotal, failedCases := part1(testCases)
	part2(subTotal, failedCases)
}
