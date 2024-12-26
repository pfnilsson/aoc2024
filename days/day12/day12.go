package day12

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

type direction int

const (
	horizontal = iota
	vertical   = iota
)

var orientations = [2]float64{-1, 1}

type edge struct {
	x         float64
	y         float64
	direction direction
	owner     shared.Point
}

func newEdge(x float64, y float64, direction direction, owner shared.Point) edge {
	return edge{x: x, y: y, direction: direction, owner: owner}
}

func (e edge) copy() edge {
	return newEdge(e.x, e.y, e.direction, e.owner)
}

func (e edge) slide(val float64) edge {
	if e.direction == horizontal {
		e.x += val
		e.owner.X += int(val)
	} else {
		e.y += val
		e.owner.Y += int(val)
	}

	return e
}

func edgeFromPoints(a shared.Point, b shared.Point) edge {
	var dir direction

	if a.X == b.X {
		dir = horizontal
	} else {
		dir = vertical
	}

	ax, ay, bx, by := float64(a.X), float64(a.Y), float64(b.X), float64(b.Y)
	return newEdge((ax+bx)/2, (ay+by)/2, dir, a)
}

func findRegions(grid shared.Grid[rune]) []shared.Set[shared.Point] {
	var current shared.Point

	regions := []shared.Set[shared.Point]{*shared.NewSet[shared.Point]()}
	seen := shared.NewSet[shared.Point]()
	stack := shared.NewStack[shared.Point](shared.NewPoint(0, 0))
	nextCandidates := shared.NewSet[shared.Point]()

	regionIndex := 0

	for {
		if stack.IsEmpty() {
			next, ok := nextCandidates.Pop()
			if !ok {
				break
			}

			if seen.Contains(next) {
				continue
			}

			stack.Push(next)
			regions = append(regions, *shared.NewSet[shared.Point]())
			regionIndex++
		}

		current, _ = stack.Pop()

		if seen.Contains(current) {
			continue
		}
		seen.Add(current)
		regions[regionIndex].Add(current)

		plant := grid.Get(current)
		for _, neighbor := range current.CardinalNeighbors() {
			if seen.Contains(neighbor) || !grid.Contains(neighbor) {
				continue
			}

			if grid.Get(neighbor) == plant {
				stack.Push(neighbor)
			} else {
				nextCandidates.Add(neighbor)
			}
		}
	}

	return regions
}

func regionPrice(region shared.Set[shared.Point]) int {
	fences := 0
	for _, plot := range region.Items() {
		for _, neighbor := range plot.CardinalNeighbors() {
			if !region.Contains(neighbor) {
				fences++
			}
		}
	}
	return fences * region.Size()
}

func findEdges(region shared.Set[shared.Point]) *shared.Set[edge] {
	edges := shared.NewSet[edge]()
	for _, plot := range region.Items() {
		for _, neighbor := range plot.CardinalNeighbors() {
			if region.Contains(neighbor) {
				continue
			}
			edges.Add(edgeFromPoints(plot, neighbor))
		}
	}
	return edges
}

func fenceSides(region shared.Set[shared.Point]) int {
	edges := findEdges(region)
	seen := shared.NewSet[edge]()
	sides := 0

	for _, e := range edges.Items() {
		if seen.Contains(e) {
			continue
		}
		seen.Add(e)
		sides++

		for _, orientation := range orientations {
			current := e
			for {
				nextEdge := current.slide(orientation)
				if !edges.Contains(nextEdge) {
					break
				}
				seen.Add(nextEdge)
				current = nextEdge
			}
		}
	}

	return sides
}

func part1(regions []shared.Set[shared.Point]) {
	totalCost := 0
	for _, region := range regions {
		totalCost += regionPrice(region)
	}

	fmt.Println("Part 1:", totalCost)
}

func part2(regions []shared.Set[shared.Point]) {
	totalCost := 0
	for _, region := range regions {
		sides := fenceSides(region)
		totalCost += sides * region.Size()
	}
	fmt.Println("Part 2:", totalCost)
}

func Run() {
	grid, err := shared.ReadFileToRuneGrid("days/day12/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	regions := findRegions(grid)

	part1(regions)
	part2(regions)
}
