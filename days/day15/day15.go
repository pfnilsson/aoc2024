package day15

import (
	"fmt"
	"log"
	"strings"

	"aoc2024/shared"
)

var moveMap = map[rune]func(shared.Point) shared.Point{
	'>': shared.Point.Right,
	'<': shared.Point.Left,
	'^': shared.Point.Up,
	'v': shared.Point.Down,
}

type wideBox struct {
	left  shared.Point
	right shared.Point
}

func newWideBox(left shared.Point, right shared.Point) wideBox {
	return wideBox{left: left, right: right}
}

func parseGrid(rawGrid []string) (shared.Grid[rune], shared.Point) {
	var startingPoint shared.Point

	grid := make([][]rune, len(rawGrid))
	for i, line := range rawGrid {
		row := []rune(line)
		grid[i] = row

		robotIndex := strings.Index(line, "@")
		if robotIndex != -1 {
			startingPoint = shared.NewPoint(robotIndex, i)
		}

	}
	return shared.NewGrid(grid), startingPoint
}

func findNewBoxPosition(grid shared.Grid[rune], move func(shared.Point) shared.Point, boxPos shared.Point) shared.Point {
	newBoxPos := move(boxPos)

	for {
		if grid.Get(newBoxPos) == '#' {
			return boxPos
		}
		if grid.Get(newBoxPos) == 'O' {
			newBoxPos = move(newBoxPos)
			continue
		}
		break
	}

	return newBoxPos
}

func findConnectedWideBoxes(grid shared.Grid[rune], move func(shared.Point) shared.Point, box wideBox) []wideBox {
	var connectedBoxes []wideBox

	nextLeft, nextRight := move(box.left), move(box.right)
	nextLeftChar, nextRightChar := grid.Get(nextLeft), grid.Get(nextRight)

	addConnectedBox := func(left, right shared.Point) {
		connectedBox := newWideBox(left, right)
		if connectedBox != box {
			connectedBoxes = append(connectedBoxes, connectedBox)
			connectedBoxes = append(connectedBoxes, findConnectedWideBoxes(grid, move, connectedBox)...)
		}
	}

	switch {
	case nextLeftChar == '[' && nextRightChar == ']':
		addConnectedBox(nextLeft, nextRight)
	case nextLeftChar == ']' && nextRightChar == '[':
		addConnectedBox(shared.NewPoint(nextLeft.X-1, nextLeft.Y), nextLeft)
		addConnectedBox(nextRight, shared.NewPoint(nextRight.X+1, nextRight.Y))
	case nextLeftChar == ']' && nextRightChar == '.':
		addConnectedBox(shared.NewPoint(nextLeft.X-1, nextLeft.Y), nextLeft)
	case nextLeftChar == '.' && nextRightChar == '[':
		addConnectedBox(nextRight, shared.NewPoint(nextRight.X+1, nextRight.Y))
	}

	return connectedBoxes
}

func checkWideBoxCanMove(grid shared.Grid[rune], move func(shared.Point) shared.Point, box wideBox) bool {
	nextLeft, nextRight := move(box.left), move(box.right)
	nextLeftChar, nextRightChar := grid.Get(nextLeft), grid.Get(nextRight)

	if nextLeftChar == '#' || nextRightChar == '#' {
		return false
	}

	return true
}

func moveWideBox(grid *shared.Grid[rune], move func(shared.Point) shared.Point, box wideBox) {
	nextLeft, nextRight := move(box.left), move(box.right)

	grid.Set(nextLeft, '[')
	grid.Set(nextRight, ']')
}

func calculateScore(grid shared.Grid[rune], char rune) int {
	tot := 0
	for y, row := range grid.Rows() {
		for x, item := range row {
			if item == char {
				tot += 100*y + x
			}
		}
	}
	return tot
}

func parseInput(rawInput [][]string) (shared.Grid[rune], shared.Point, string) {
	rawGrid := rawInput[0]
	grid, pos := parseGrid(rawGrid)
	moves := strings.Join(rawInput[1], "")
	return grid, pos, moves
}

func widenGrid(grid shared.Grid[rune]) shared.Grid[rune] {
	newWidth := len(grid.Rows()[0]) * 2
	newHeight := len(grid.Rows())
	newGrid := shared.NewEmptyGrid(newWidth, newHeight, '.')

	for y, row := range grid.Rows() {
		for x, item := range row {
			newPoint1 := shared.NewPoint(2*x, y)
			newPoint2 := shared.NewPoint(2*x+1, y)

			switch item {
			case 'O':
				newGrid.Set(newPoint1, '[')
				newGrid.Set(newPoint2, ']')
			case '@':
				newGrid.Set(newPoint1, '@')
			case '#':
				newGrid.Set(newPoint1, '#')
				newGrid.Set(newPoint2, '#')
			default:
				continue
			}
		}
	}

	return newGrid
}

func clearCurrentPositions(grid *shared.Grid[rune], boxes []wideBox) {
	for _, box := range boxes {
		grid.Set(box.left, '.')
		grid.Set(box.right, '.')
	}
}

func moveWideBoxes(grid *shared.Grid[rune], move func(shared.Point) shared.Point, boxes []wideBox) {
	for _, box := range boxes {
		moveWideBox(grid, move, box)
	}
}

func createNextWideBox(next shared.Point, char rune) wideBox {
	if char == '[' {
		return newWideBox(next, shared.NewPoint(next.X+1, next.Y))
	}
	return newWideBox(shared.NewPoint(next.X-1, next.Y), next)
}

func allBoxesCanMove(grid shared.Grid[rune], move func(shared.Point) shared.Point, boxes []wideBox) bool {
	for _, box := range boxes {
		if !checkWideBoxCanMove(grid, move, box) {
			return false
		}
	}
	return true
}

func part1(grid shared.Grid[rune], pos shared.Point, moves string) {
	var next shared.Point
	var newBoxPos shared.Point
	var nextChar rune

	for _, move := range moves {
		next = moveMap[move](pos)
		nextChar = grid.Get(next)

		if nextChar == '#' {
			continue
		}

		if nextChar == 'O' {
			newBoxPos = findNewBoxPosition(grid, moveMap[move], next)
			if newBoxPos == next {
				continue
			}

			grid.Set(next, '.')
			grid.Set(newBoxPos, 'O')
		}

		pos = next
	}

	score := calculateScore(grid, 'O')
	fmt.Println("Part 1:", score)
}

func part2(grid shared.Grid[rune], pos shared.Point, moves string) {
	for _, move := range moves {
		next := moveMap[move](pos)
		nextChar := grid.Get(next)

		if nextChar == '#' {
			continue
		}

		if nextChar == '[' || nextChar == ']' {
			box := createNextWideBox(next, nextChar)

			connectedBoxes := findConnectedWideBoxes(grid, moveMap[move], box)
			connectedBoxes = append(connectedBoxes, box)

			if allBoxesCanMove(grid, moveMap[move], connectedBoxes) {
				clearCurrentPositions(&grid, connectedBoxes)
				moveWideBoxes(&grid, moveMap[move], connectedBoxes)
			} else {
				continue
			}
		}

		pos = next
	}

	fmt.Println("Part 2:", calculateScore(grid, '['))
}

func Run() {
	rawInput, err := shared.ReadFileByBlankLine("days/day15/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	grid, pos, moves := parseInput(rawInput)
	wideGrid := widenGrid(grid)
	wideStartingPos := shared.NewPoint(2*pos.X, pos.Y)

	part1(grid.Clone(), pos, moves)
	part2(wideGrid, wideStartingPos, moves)
}
