package day04

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

var letterIndex = map[int]rune{
	0: 'M',
	1: 'A',
	2: 'S',
}

func findUpwards(pt shared.Point, length int) []shared.Point {
	var up []shared.Point
	for i := 1; i < length+1; i++ {
		up = append(up, shared.Point{X: pt.X, Y: pt.Y - i})
	}
	return up
}

func findDownwards(pt shared.Point, length int) []shared.Point {
	var down []shared.Point
	for i := 1; i < length+1; i++ {
		down = append(down, shared.Point{X: pt.X, Y: pt.Y + i})
	}
	return down
}

func findLeft(pt shared.Point, length int) []shared.Point {
	var left []shared.Point
	for i := 1; i < length+1; i++ {
		left = append(left, shared.Point{X: pt.X - i, Y: pt.Y})
	}
	return left
}

func findRight(pt shared.Point, length int) []shared.Point {
	var right []shared.Point
	for i := 1; i < length+1; i++ {
		right = append(right, shared.Point{X: pt.X + i, Y: pt.Y})
	}
	return right
}

func findDiagonalUpLeft(pt shared.Point, length int) []shared.Point {
	var diag []shared.Point
	for i := 1; i < length+1; i++ {
		diag = append(diag, shared.Point{X: pt.X - i, Y: pt.Y - i})
	}
	return diag
}

func findDiagonalUpRight(pt shared.Point, length int) []shared.Point {
	var diag []shared.Point
	for i := 1; i < length+1; i++ {
		diag = append(diag, shared.Point{X: pt.X + i, Y: pt.Y - i})
	}
	return diag
}

func findDiagonalDownLeft(pt shared.Point, length int) []shared.Point {
	var diag []shared.Point
	for i := 1; i < length+1; i++ {
		diag = append(diag, shared.Point{X: pt.X - i, Y: pt.Y + i})
	}
	return diag
}

func findDiagonalDownRight(pt shared.Point, length int) []shared.Point {
	var diag []shared.Point
	for i := 1; i < length+1; i++ {
		diag = append(diag, shared.Point{X: pt.X + i, Y: pt.Y + i})
	}
	return diag
}

func findCandidatesPart1(pt shared.Point) [][]shared.Point {
	length := 3
	return [][]shared.Point{
		findUpwards(pt, length),
		findDownwards(pt, length),
		findLeft(pt, length),
		findRight(pt, length),
		findDiagonalUpLeft(pt, length),
		findDiagonalUpRight(pt, length),
		findDiagonalDownLeft(pt, length),
		findDiagonalDownRight(pt, length)}
}

func findCandidatesPart2(pt shared.Point) []shared.Point {
	length := 1
	return []shared.Point{
		findDiagonalUpLeft(pt, length)[0],
		findDiagonalUpRight(pt, length)[0],
		findDiagonalDownLeft(pt, length)[0],
		findDiagonalDownRight(pt, length)[0],
	}
}

func checkCandidatePart1(grid shared.Grid[rune], candidate []shared.Point) bool {
	for i, pt := range candidate {
		if !grid.Contains(pt) {
			return false
		}
		if letterIndex[i] != grid.Get(pt) {
			return false
		}
	}
	return true
}

func checkCandidatePart2(grid shared.Grid[rune], candidate []shared.Point) bool {
	for _, pt := range candidate {
		if !grid.Contains(pt) {
			return false
		}
	}

	upLeft := candidate[0]
	upRight := candidate[1]
	downLeft := candidate[2]
	downRight := candidate[3]

	leftDiag := string(grid.Get(upLeft)) + string(grid.Get(downRight))
	if leftDiag != "MS" && leftDiag != "SM" {
		return false
	}

	rightDiag := string(grid.Get(upRight)) + string(grid.Get(downLeft))
	if rightDiag != "MS" && rightDiag != "SM" {
		return false
	}
	return true
}

func part1(grid shared.Grid[rune]) {
	tot := 0
	for y, row := range grid.Rows() {
		for x, letter := range row {
			if letter != 'X' {
				continue
			}

			pt := shared.Point{X: x, Y: y}
			candidates := findCandidatesPart1(pt)
			for _, candidate := range candidates {
				if checkCandidatePart1(grid, candidate) {
					tot += 1
				}
			}
		}
	}
	fmt.Println("Part 1:", tot)
}

func part2(grid shared.Grid[rune]) {
	tot := 0
	for y, row := range grid.Rows() {
		for x, letter := range row {
			if letter != 'A' {
				continue
			}

			pt := shared.Point{X: x, Y: y}
			candidates := findCandidatesPart2(pt)
			if checkCandidatePart2(grid, candidates) {
				tot += 1
			}
		}
	}
	fmt.Println("Part 2:", tot)
}

func Run() {
	grid, err := shared.ReadFileToRuneGrid("days/day04/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(grid)
	part2(grid)
}
