package day20

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

const minTimeSave = 100

func getPath(grid shared.Grid[rune], startingPoint shared.Point, goal shared.Point) (map[shared.Point]int, []shared.Point) {
	var path []shared.Point
	pathMap := make(map[shared.Point]int)

	curr := startingPoint
	for i := 0; ; i++ {
		pathMap[curr] = i
		path = append(path, curr)

		if curr == goal {
			break
		}

		for _, neighbor := range curr.CardinalNeighbors() {
			_, seen := pathMap[neighbor]
			if seen || !grid.Contains(neighbor) || grid.Get(neighbor) == '#' {
				continue
			}
			curr = neighbor
		}
	}

	return pathMap, path
}

func getCheatSpots(wall shared.Point, grid shared.Grid[rune]) []shared.Point {
	var cheatSpots []shared.Point

	for _, cheatSpot := range wall.CardinalNeighbors() {
		if !grid.Contains(cheatSpot) || grid.Get(cheatSpot) == '#' {
			continue
		}
		cheatSpots = append(cheatSpots, cheatSpot)
	}
	return cheatSpots
}

func part1(grid shared.Grid[rune], path map[shared.Point]int) {
	cheatCounts := 0

	for y, row := range grid.Rows() {
		for x, char := range row {
			if char != '#' {
				continue
			}

			wall := shared.NewPoint(x, y)
			cheatSpots := getCheatSpots(wall, grid)
			for _, cheatStart := range cheatSpots {
				for _, cheatEnd := range cheatSpots {
					timeSave := path[cheatEnd] - path[cheatStart] - 2

					if timeSave >= minTimeSave {
						cheatCounts++
					}
				}
			}
		}
	}
	fmt.Println("Part 1:", cheatCounts)
}

func part2(path []shared.Point) {
	lastCandidate := len(path) - 1 - minTimeSave
	cheatCount := 0
	for i, cheatStart := range path[:lastCandidate] {
		for j, cheatEnd := range path[i:] {
			distance := shared.ManhattanDistance(cheatStart, cheatEnd)
			if distance > 20 {
				continue
			}

			timeSave := j - distance
			if timeSave >= minTimeSave {
				cheatCount++
			}
		}
	}
	fmt.Println("Part 2:", cheatCount)
}

func Run() {
	grid, startingPoint, goal, err := shared.ReadFileToRuneGridWithStartingPointAndGoal(
		"days/day20/input.txt", 'S', 'E',
	)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	pathMap, path := getPath(grid, startingPoint, goal)

	part1(grid, pathMap)
	part2(path)
}
