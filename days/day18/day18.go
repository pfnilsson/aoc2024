package day18

import (
	"fmt"
	"log"

	"aoc2024/shared"
)

func findShortestPath(graph map[shared.Point][]shared.Point, start shared.Point, goal shared.Point) (shared.Set[shared.Point], error) {
	var current shared.Point

	distances := make(map[shared.Point]int)
	parent := make(map[shared.Point]shared.Point)
	distances[start] = 0
	queue := []shared.Point{start}

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, neighbor := range graph[current] {
			if _, ok := distances[neighbor]; ok {
				continue
			}

			distances[neighbor] = distances[current] + 1
			parent[neighbor] = current
			queue = append(queue, neighbor)

			if neighbor == goal {
				return reconstructPath(parent, start, goal), nil
			}
		}
	}

	return *shared.NewSet[shared.Point](), fmt.Errorf("no path found")
}

func reconstructPath(parent map[shared.Point]shared.Point, start, end shared.Point) shared.Set[shared.Point] {
	path := shared.NewSet(start)
	for current := end; current != start; current = parent[current] {
		path.Add(current)
	}
	return *path
}

func graphFromGrid(grid shared.Grid[rune]) map[shared.Point][]shared.Point {
	graph := make(map[shared.Point][]shared.Point)
	for y, row := range grid.Rows() {
		for x, char := range row {
			if char == '#' {
				continue
			}
			current := shared.NewPoint(x, y)
			for _, neighbor := range current.CardinalNeighbors() {
				if !grid.Contains(neighbor) || grid.Get(neighbor) == '#' {
					continue
				}
				graph[current] = append(graph[current], neighbor)
			}
		}
	}

	return graph
}

func createInitialGrid(lines [][]int) shared.Grid[rune] {
	grid := shared.NewEmptyGrid(71, 71, '.')
	for i, line := range lines {
		if i == 1024 {
			break
		}
		pt := shared.NewPoint(line[0], line[1])
		grid.Set(pt, '#')
	}
	return grid
}

func intialSetup(lines [][]int) (shared.Grid[rune], map[shared.Point][]shared.Point, shared.Point, shared.Point) {
	grid := createInitialGrid(lines)
	graph := graphFromGrid(grid)
	start := shared.NewPoint(0, 0)
	end := shared.NewPoint(70, 70)
	return grid, graph, start, end
}

func part1(path shared.Set[shared.Point]) {
	fmt.Println("Part 1:", path.Size()-1)
}

func part2(
	lines [][]int,
	grid shared.Grid[rune],
	start shared.Point, end shared.Point,
	path shared.Set[shared.Point],
) {
	for i := 1024; i < len(lines); i++ {
		x, y := lines[i][0], lines[i][1]
		pt := shared.NewPoint(x, y)
		grid.Set(pt, '#')
		if !path.Contains(pt) {
			continue
		}

		graph := graphFromGrid(grid)

		var err error
		path, err = findShortestPath(graph, start, end)
		if err != nil {
			fmt.Printf("Part 2: %v,%v\n", x, y)
			return
		}
	}
}

func Run() {
	lines, err := shared.ReadFileByLineToSplitInts("days/day18/input.txt", ",")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	grid, graph, start, end := intialSetup(lines)
	path, err := findShortestPath(graph, start, end)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(path)
	part2(lines, grid, start, end, path)
}
