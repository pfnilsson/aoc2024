package day06

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"sync"
)

type state struct {
	position  shared.Point
	direction string
}

func newState(position shared.Point, direction string) state {
	return state{position, direction}
}

func (s state) clone() state {
	return newState(s.position, s.direction)
}

var directionOrder = map[string]string{
	"up":    "right",
	"right": "down",
	"down":  "left",
	"left":  "up",
}

func getRoute(startingPoint shared.Point, grid shared.Grid[rune]) (route *shared.Set[shared.Point], loop bool) {

	currentState := newState(startingPoint, "up")
	route = shared.NewSet[shared.Point](currentState.position)
	seenStates := shared.NewSet[state]()
	var nextPosition shared.Point

	for {
		if seenStates.Contains(currentState) {
			return route, true
		}
		seenStates.Add(currentState.clone())

		switch currentState.direction {
		case "up":
			nextPosition = currentState.position.Up()
		case "down":
			nextPosition = currentState.position.Down()
		case "left":
			nextPosition = currentState.position.Left()
		case "right":
			nextPosition = currentState.position.Right()
		}

		if !grid.Contains(nextPosition) {
			return route, false
		}

		if grid.Get(nextPosition) == '#' {
			currentState.direction = directionOrder[currentState.direction]
			continue
		}

		currentState.position = nextPosition
		route.Add(currentState.position)
	}
}

func part1(route *shared.Set[shared.Point]) {
	fmt.Println("Part 1:", route.Size())
}

func part2(startingPoint shared.Point, route *shared.Set[shared.Point], grid shared.Grid[rune]) {
	candidatePoints := make([]shared.Point, 0, route.Size())

	for _, point := range route.Items() {
		if grid.Get(point) != '#' {
			candidatePoints = append(candidatePoints, point)
		}
	}

	tot := 0
	candidateCount := len(candidatePoints)

	results := make(chan bool, candidateCount)

	var wg sync.WaitGroup
	wg.Add(candidateCount)

	for _, point := range candidatePoints {
		go func(p shared.Point) {
			defer wg.Done()

			clonedGrid := grid.Clone()
			clonedGrid.Set(p, '#')
			_, loop := getRoute(startingPoint, clonedGrid)

			results <- loop
		}(point)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for loop := range results {
		if loop {
			tot++
		}
	}

	fmt.Println("Part 2:", tot)
}

func Run() {
	grid, startingPoint, err := shared.ReadFileToRuneGridWithStartingPoint("days/day06/input.txt", '^')
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	route, _ := getRoute(startingPoint, grid)

	part1(route)
	part2(startingPoint, route, grid)
}
