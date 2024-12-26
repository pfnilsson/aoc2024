package day16

import (
	"aoc2024/shared"
	"container/heap"
	"fmt"
	"log"
	"math"
)

var directionMap = map[direction]int{
	{0, 1}:  0,
	{1, 0}:  1,
	{0, -1}: 2,
	{-1, 0}: 3,
}

type direction struct {
	x int
	y int
}

func newDirection(x, y int) direction {
	return direction{x: x, y: y}
}

type node struct {
	point     shared.Point
	direction direction
}

func newNode(point shared.Point, direction direction) node {
	return node{point: point, direction: direction}
}

type edge struct {
	destination node
	cost        int
}

func newEdge(destination node, cost int) edge {
	return edge{destination: destination, cost: cost}
}

func turnCost(from, to direction) int {
	turnSteps := (directionMap[to] - directionMap[from]) % 4
	if turnSteps < 0 {
		turnSteps += 4
	}

	if turnSteps == 0 {
		return 1
	}

	if turnSteps == 2 {
		return 2001
	}

	return 1001
}

func graphFromGrid(grid shared.Grid[rune]) map[node][]edge {
	graph := make(map[node][]edge)
	for y, row := range grid.Rows() {
		for x, char := range row {
			if char == '#' {
				continue
			}
			current := shared.NewPoint(x, y)
			updateGraphWithPos(grid, current, graph)
		}
	}

	return graph
}

func updateGraphWithPos(grid shared.Grid[rune], current shared.Point, graph map[node][]edge) {
	for currDir := range directionMap {
		currNode := newNode(current, currDir)

		for nextDir := range directionMap {
			neighbor := shared.NewPoint(current.X+nextDir.x, current.Y+nextDir.y)

			if !grid.Contains(neighbor) || grid.Get(neighbor) == '#' {
				continue
			}

			cost := turnCost(currDir, nextDir)
			neighborNode := newNode(neighbor, nextDir)
			currEdge := newEdge(neighborNode, cost)
			graph[currNode] = append(graph[currNode], currEdge)
		}
	}
}

func findGoal(grid shared.Grid[rune]) shared.Point {
	for y, row := range grid.Rows() {
		for x, char := range row {
			if char == 'E' {
				return shared.NewPoint(x, y)
			}
		}
	}
	return shared.Point{}
}

func dijkstra(graph map[node][]edge, start node, goal shared.Point) (int, map[node][]node, map[node]int) {
	distances := make(map[node]int)
	predecessors := make(map[node][]node)

	pq := shared.NewPriorityQueue[node]()
	heap.Init(&pq)

	for nd := range graph {
		distances[nd] = math.MaxInt
		predecessors[nd] = nil
	}

	distances[start] = 0
	startEntry := shared.NewPriorityQueueEntry[node](start, 0)
	heap.Push(&pq, &startEntry)

	shortestDistance := math.MaxInt

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*shared.PriorityQueueEntry[node])
		if distances[current.Value] < current.Priority {
			continue
		}

		if current.Value.point == goal {
			if distances[current.Value] < shortestDistance {
				shortestDistance = distances[current.Value]
			}
		}

		for _, neighbor := range graph[current.Value] {
			destination := neighbor.destination
			alternative := distances[current.Value] + neighbor.cost

			if alternative < distances[destination] {
				distances[destination] = alternative
				predecessors[destination] = []node{current.Value}
				newEntry := shared.NewPriorityQueueEntry[node](destination, alternative)
				heap.Push(&pq, &newEntry)
			} else if alternative == distances[destination] {
				predecessors[destination] = append(predecessors[destination], current.Value)
			}
		}
	}

	return shortestDistance, predecessors, distances
}

func reconstructPaths(current node, currentPath []shared.Point, predecessors map[node][]node, start node, allPaths *[][]shared.Point) {
	currentPath = append([]shared.Point{current.point}, currentPath...)

	if current == start {
		*allPaths = append(*allPaths, currentPath)
		return
	}

	for _, pred := range predecessors[current] {
		reconstructPaths(pred, currentPath, predecessors, start, allPaths)
	}
}

func GetAllShortestPaths(predecessors map[node][]node, start, goal node) [][]shared.Point {
	var allPaths [][]shared.Point

	reconstructPaths(goal, nil, predecessors, start, &allPaths)
	return allPaths
}

func GetPointsOnPaths(paths [][]shared.Point) shared.Set[shared.Point] {
	pointsSet := shared.NewSet[shared.Point]()
	for _, path := range paths {
		for _, p := range path {
			pointsSet.Add(p)
		}
	}
	return *pointsSet
}

func getGoalNodes(distances map[node]int, goal shared.Point, shortestDistance int) []node {
	var goalNodes []node
	for n, d := range distances {
		if n.point == goal && d == shortestDistance {
			goalNodes = append(goalNodes, n)
		}
	}
	return goalNodes
}

func getAllShortestPathsAllGoals(goalNodes []node, predecessors map[node][]node, startNode node) [][]shared.Point {
	var allPaths [][]shared.Point
	for _, g := range goalNodes {
		paths := GetAllShortestPaths(predecessors, startNode, g)
		allPaths = append(allPaths, paths...)
	}
	return allPaths
}

func part1(shortestDistance int) {
	fmt.Println("Part 1:", shortestDistance)
}

func part2(shortestDistance int, predecessors map[node][]node, distances map[node]int, goal shared.Point, startNode node) {
	goalNodes := getGoalNodes(distances, goal, shortestDistance)
	allPaths := getAllShortestPathsAllGoals(goalNodes, predecessors, startNode)
	allPoints := GetPointsOnPaths(allPaths)

	fmt.Println("Part 2:", allPoints.Size())
}

func Run() {
	grid, start, err := shared.ReadFileToRuneGridWithStartingPoint("days/day16/input.txt", 'S')
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	goal := findGoal(grid)
	graph := graphFromGrid(grid)
	startNode := newNode(start, newDirection(1, 0))

	shortestDistance, predecessors, distances := dijkstra(graph, startNode, goal)

	part1(shortestDistance)
	part2(shortestDistance, predecessors, distances, goal, startNode)

}
