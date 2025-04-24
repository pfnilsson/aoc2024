package day08

import (
	"fmt"
	"log"

	"aoc2024/shared"
)

type antennaPair struct {
	a shared.Point
	b shared.Point
}

func newAntennaPair(a shared.Point, b shared.Point) antennaPair {
	return antennaPair{a: a, b: b}
}

func (p *antennaPair) antinodes(grid shared.Grid[rune]) []shared.Point {
	var antinodes []shared.Point

	addPointIfValid := func(x, y int) {
		point := shared.NewPoint(x, y)
		if grid.Contains(point) {
			antinodes = append(antinodes, point)
		}
	}

	xDist := p.a.X - p.b.X
	yDist := p.a.Y - p.b.Y

	addPointIfValid(p.a.X+xDist, p.a.Y+yDist)
	addPointIfValid(p.b.X-xDist, p.b.Y-yDist)

	return antinodes
}

func (p *antennaPair) antinodesWithResonantHarmonics(grid shared.Grid[rune]) []shared.Point {
	var antinodes []shared.Point

	xDist := p.a.X - p.b.X
	yDist := p.a.Y - p.b.Y

	addPointsInDirection := func(start shared.Point, step int) {
		for i := 0; ; i++ {
			nextPoint := shared.NewPoint(start.X+xDist*i*step, start.Y+yDist*i*step)
			if !grid.Contains(nextPoint) {
				break
			}
			antinodes = append(antinodes, nextPoint)
		}
	}

	addPointsInDirection(p.a, 1)
	addPointsInDirection(p.a, -1)

	return antinodes
}

func getAntennaPairs(antennas map[rune][]shared.Point) []antennaPair {
	var pairs []antennaPair

	for _, locations := range antennas {
		for i := range locations {
			for j := i + 1; j < len(locations); j++ {
				pairs = append(pairs, newAntennaPair(locations[i], locations[j]))
			}
		}
	}
	return pairs
}

func findAntennas(grid shared.Grid[rune]) map[rune][]shared.Point {
	antennas := make(map[rune][]shared.Point)

	for y, row := range grid.Rows() {
		for x, char := range row {
			if char == '.' {
				continue
			}
			antennas[char] = append(antennas[char], shared.NewPoint(x, y))
		}
	}

	return antennas
}

func part1(grid shared.Grid[rune], pairs []antennaPair) {
	antinodeLocations := shared.NewSet[shared.Point]()

	for _, pair := range pairs {
		antinodes := pair.antinodes(grid)
		for _, antinode := range antinodes {
			antinodeLocations.Add(antinode)
		}
	}

	fmt.Println("Part 1:", antinodeLocations.Size())
}

func part2(grid shared.Grid[rune], pairs []antennaPair) {
	antinodeLocations := shared.NewSet[shared.Point]()

	for _, pair := range pairs {
		antinodes := pair.antinodesWithResonantHarmonics(grid)
		for _, antinode := range antinodes {
			antinodeLocations.Add(antinode)
		}
	}

	fmt.Println("Part 2:", antinodeLocations.Size())
}

func Run() {
	grid, err := shared.ReadFileToRuneGrid("days/day08/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	antennas := findAntennas(grid)
	pairs := getAntennaPairs(antennas)

	part1(grid, pairs)
	part2(grid, pairs)
}
