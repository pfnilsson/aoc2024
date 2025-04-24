package day03

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"aoc2024/shared"
)

func strMul(xStr string, yStr string) (int, error) {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return 0, err
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return 0, err
	}

	return x * y, nil
}

func part1(memory string) error {
	re := regexp.MustCompile(`mul\((\d{1,3}),\s*(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(memory, -1)

	tot := 0
	for _, match := range matches {
		xStr := match[1]
		yStr := match[2]

		res, err := strMul(xStr, yStr)
		if err != nil {
			return err
		}

		tot += res
	}
	fmt.Println("Part 1:", tot)
	return nil
}

func part2(memory string) error {
	pattern := `(do|don't|mul)\((\d{1,3})?,?(\d{1,3})?\)`
	regex := regexp.MustCompile(pattern)

	matches := regex.FindAllStringSubmatchIndex(memory, -1)
	tot := 0
	enabled := true
	for _, match := range matches {
		funcName := memory[match[2]:match[3]]

		if funcName == "do" {
			enabled = true
		} else if funcName == "don't" {
			enabled = false
		} else if funcName == "mul" {
			if !enabled {
				continue
			}

			xStart, xEnd := match[4], match[5]
			yStart, yEnd := match[6], match[7]

			xStr := memory[xStart:xEnd]
			yStr := memory[yStart:yEnd]

			res, err := strMul(xStr, yStr)
			if err != nil {
				return err
			}

			tot += res

		}
	}
	fmt.Println("Part 2:", tot)
	return nil
}

func Run() {
	memory, err := shared.ReadFileToString("days/day03/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = part1(memory)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = part2(memory)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
