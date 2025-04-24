package day13

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"aoc2024/shared"
)

const ConversionError = 10_000_000_000_000

type vector struct {
	x int
	y int
}

func newVector(x int, y int) vector {
	return vector{x: x, y: y}
}

func (v *vector) addToAll(val int) vector {
	return newVector(v.x+val, v.y+val)
}

func determinant(v1 vector, v2 vector) int {
	return v1.x*v2.y - v1.y*v2.x
}

func solveLinearSystem(a vector, b vector, solution vector) (int, int, error) {
	det := determinant(a, b)
	if det == 0 {
		return 0, 0, fmt.Errorf("the system has no unique solution, determinant is zero")
	}

	numeratorA := b.y*solution.x - b.x*solution.y
	numeratorB := -a.y*solution.x + a.x*solution.y

	if numeratorA%det != 0 || numeratorB%det != 0 {
		return 0, 0, fmt.Errorf("non-integer solution")
	}

	aSolution := numeratorA / det
	bSolution := numeratorB / det

	return aSolution, bSolution, nil
}

type machine struct {
	a             vector
	b             vector
	prizeLocation vector
}

func newMachine(a vector, b vector, prizeLocation vector) machine {
	return machine{a: a, b: b, prizeLocation: prizeLocation}
}

func parseInput(rawInput [][]string) ([]machine, error) {
	machines := make([]machine, len(rawInput))
	for i, config := range rawInput {
		m, err := parseMachineConfig(config)
		if err != nil {
			return nil, err
		}
		machines[i] = m
	}
	return machines, nil
}

func parseMachineConfig(machineConfig []string) (machine, error) {
	re := regexp.MustCompile(`(?:X=|X\+)(\d+),\s?(?:Y=|Y\+)(\d+)`)
	xs, ys := make([]int, 3), make([]int, 3)

	for i, line := range machineConfig {
		match := re.FindStringSubmatch(line)
		if match == nil {
			return machine{}, fmt.Errorf("invalid line format: %s", line)
		}

		x, err := strconv.Atoi(match[1])
		if err != nil {
			return machine{}, fmt.Errorf("invalid X value in line %d: %v", i, err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			return machine{}, fmt.Errorf("invalid Y value in line %d: %v", i, err)
		}

		xs[i], ys[i] = x, y
	}

	buttonA := newVector(xs[0], ys[0])
	buttonB := newVector(xs[1], ys[1])
	prizeLocation := newVector(xs[2], ys[2])

	return newMachine(buttonA, buttonB, prizeLocation), nil
}

func findMachineCost(m machine, errorCorrection int) (int, error) {
	a, b, err := solveLinearSystem(m.a, m.b, m.prizeLocation.addToAll(errorCorrection))
	if err != nil {
		return 0, err
	}

	return 3*a + b, nil
}

func part1(machines []machine) {
	tot := 0
	for _, m := range machines {
		cost, err := findMachineCost(m, 0)
		if err != nil {
			continue
		}
		tot += cost
	}
	fmt.Println("Part 1:", tot)
}

func part2(machines []machine) {
	tot := 0
	for _, m := range machines {
		cost, err := findMachineCost(m, ConversionError)
		if err != nil {
			continue
		}
		tot += cost
	}
	fmt.Println("Part 2:", tot)
}

func Run() {
	rawInput, err := shared.ReadFileByBlankLine("days/day13/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	machines, err := parseInput(rawInput)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(machines)
	part2(machines)
}
