package day21

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var (
	initialKey = 'A'
	left       = '<'
	right      = '>'
	up         = '^'
	down       = 'v'
	empty      = ' '
	activate   = 'A'
)

var numPad = map[rune]shared.Point{
	'7': {0, 0}, '8': {1, 0}, '9': {2, 0},
	'4': {0, 1}, '5': {1, 1}, '6': {2, 1},
	'1': {0, 2}, '2': {1, 2}, '3': {2, 2},
	' ': {0, 3}, '0': {1, 3}, 'A': {2, 3}}

var dirPad = map[rune]shared.Point{
	' ': {0, 0}, '^': {1, 0}, 'A': {2, 0},
	'<': {0, 1}, 'v': {1, 1}, '>': {2, 1},
}

func getNumericValue(code string) (int, error) {
	numPart := code[:len(code)-1]
	val, err := strconv.Atoi(numPart)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func buildSteps(dist int, posDir rune, negDir rune) string {
	if dist > 0 {
		return strings.Repeat(string(posDir), dist)
	} else if dist < 0 {
		return strings.Repeat(string(negDir), -dist)
	} else {
		return ""
	}
}

func getFromToMissing(start rune, end rune) (shared.Point, shared.Point, shared.Point) {
	from, startOk := dirPad[start]
	to, endOk := dirPad[end]
	missing := dirPad[empty]

	if !startOk || !endOk {
		from, _ = numPad[start]
		to, _ = numPad[end]
		missing = numPad[empty]
	}
	return from, to, missing
}

func getTurnPoints(from shared.Point, to shared.Point) (shared.Point, shared.Point) {
	horTurnPoint := shared.NewPoint(to.X, from.Y)
	verTurnPoint := shared.NewPoint(from.X, to.Y)
	return horTurnPoint, verTurnPoint
}

func verFirst(from shared.Point, to shared.Point, missing shared.Point, xDist int) bool {
	horTurnPoint, verTurnPoint := getTurnPoints(from, to)

	ret := xDist > 0
	if horTurnPoint == missing {
		ret = true
	} else if verTurnPoint == missing {
		ret = false
	}
	return ret
}

func getXYDistances(from shared.Point, to shared.Point) (int, int) {
	xDist := to.X - from.X
	yDist := to.Y - from.Y
	return xDist, yDist
}

func findShortestSequence(start rune, end rune) string {
	from, to, missing := getFromToMissing(start, end)
	xDist, yDist := getXYDistances(from, to)

	horizontalSteps := buildSteps(xDist, right, left)
	verticalSteps := buildSteps(yDist, down, up)

	var sb strings.Builder
	if verFirst(from, to, missing, xDist) {
		sb.WriteString(verticalSteps)
		sb.WriteString(horizontalSteps)
	} else {
		sb.WriteString(horizontalSteps)
		sb.WriteString(verticalSteps)
	}

	sb.WriteRune(activate)
	return sb.String()
}

func generateCacheKey(code string, depth int) string {
	return fmt.Sprintf("%s|%d", code, depth)
}

func getLength(code string, depth int, cache map[string]int) int {
	cacheKey := generateCacheKey(code, depth)

	if val, ok := cache[cacheKey]; ok {
		return val
	}

	var length int
	if depth == 0 {
		length = len(code)
		cache[cacheKey] = length
		return length
	}

	start := initialKey
	for _, char := range code {
		seq := findShortestSequence(start, char)
		length += getLength(seq, depth-1, cache)
		start = char
	}

	cache[cacheKey] = length
	return length
}

func complexitySum(codes []string, depth int, cache map[string]int) (int, error) {
	total := 0
	for _, code := range codes {
		numVal, err := getNumericValue(code)
		if err != nil {
			return 0, err
		}
		total += getLength(code, depth, cache) * numVal
	}
	return total, nil
}

func part1(codes []string, cache map[string]int) error {
	tot, err := complexitySum(codes, 3, cache)
	if err != nil {
		return err
	}

	fmt.Println("Part 1:", tot)
	return nil
}

func part2(codes []string, cache map[string]int) error {
	tot, err := complexitySum(codes, 26, cache)
	if err != nil {
		return err
	}

	fmt.Println("Part 2:", tot)
	return nil
}

func Run() {
	codes, err := shared.ReadFileByLine("days/day21/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	cache := make(map[string]int)

	err = part1(codes, cache)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	err = part2(codes, cache)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
}
