package day10

import (
	"fmt"
	"log"

	"aoc2024/shared"
)

const top = 9

func findTrailheads(grid shared.Grid[int]) []shared.Point {
	var trailheads []shared.Point

	for y, row := range grid.Rows() {
		for x, height := range row {
			if height == 0 {
				trailheads = append(trailheads, shared.NewPoint(x, y))
			}
		}
	}

	return trailheads
}

func possibleTrails(grid shared.Grid[int], trailhead shared.Point) []shared.Point {
	var reachableTrailheads []shared.Point
	currentHeight := grid.Get(trailhead)

	for _, neighbor := range trailhead.CardinalNeighbors() {
		if !grid.Contains(neighbor) {
			continue
		}

		neighborHeight := grid.Get(neighbor)
		if neighborHeight == currentHeight+1 {
			if neighborHeight == top {
				reachableTrailheads = append(reachableTrailheads, neighbor)
			} else {
				reachableTrailheads = append(reachableTrailheads, possibleTrails(grid, neighbor)...)
			}
		}
	}

	return reachableTrailheads
}

func part1(trailheads []shared.Point, grid shared.Grid[int]) {
	tot := 0
	for _, trailhead := range trailheads {
		trails := possibleTrails(grid, trailhead)
		tot += len(shared.UniqueSlice(trails))
	}
	fmt.Println("Part 1:", tot)
}

func part2(trailheads []shared.Point, grid shared.Grid[int]) {
	tot := 0
	for _, trailhead := range trailheads {
		tot += len(possibleTrails(grid, trailhead))
	}
	fmt.Println("Part 2:", tot)
}

func Run() {
	grid, err := shared.ReadFileToIntGrid("days/day10/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	trailheads := findTrailheads(grid)

	part1(trailheads, grid)
	part2(trailheads, grid)
}
