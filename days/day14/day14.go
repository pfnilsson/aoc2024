package day14

import (
	"aoc2024/shared"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const roomWidth = 101
const roomHeight = 103

type roomSize struct {
	width  int
	height int
}

func newRoomSize(width int, height int) roomSize {
	return roomSize{width, height}
}

type velocity struct {
	x int
	y int
}

func newVelocity(x int, y int) velocity {
	return velocity{x, y}
}

type robot struct {
	position shared.Point
	velocity velocity
}

func newRobot(position shared.Point, velocity velocity) robot {
	return robot{position, velocity}
}

func (r robot) positionAtTime(t int, rs roomSize) shared.Point {
	newX := (r.position.X + r.velocity.x*t) % rs.width
	if newX < 0 {
		newX += rs.width
	}
	newY := (r.position.Y + r.velocity.y*t) % rs.height
	if newY < 0 {
		newY += rs.height
	}
	return shared.NewPoint(newX, newY)
}

func (r robot) quadrantAtTime(t int, rs roomSize) int {
	pos := r.positionAtTime(t, rs)
	horizontalMid := (rs.width - 1) / 2
	verticalMid := (rs.height - 1) / 2

	switch {
	case pos.X < horizontalMid && pos.Y < verticalMid:
		return 1
	case pos.X < horizontalMid && pos.Y > verticalMid:
		return 3
	case pos.X > horizontalMid && pos.Y < verticalMid:
		return 2
	case pos.X > horizontalMid && pos.Y > verticalMid:
		return 4
	default:
		return 0
	}
}

func parseRobot(robotStr string) (robot, error) {
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)
	matches := re.FindStringSubmatch(robotStr)
	intMatches := make([]int, 4)

	for i := 0; i < 4; i++ {
		intMatch, err := strconv.Atoi(matches[i+1])
		if err != nil {
			return robot{}, fmt.Errorf("converting regex match to int: %w", err)
		}

		intMatches[i] = intMatch
	}
	pos := shared.NewPoint(intMatches[0], intMatches[1])
	vel := newVelocity(intMatches[2], intMatches[3])
	return newRobot(pos, vel), nil
}

func parseInput(input []string) ([]robot, error) {
	robots := make([]robot, len(input))

	for i, line := range input {
		bot, err := parseRobot(line)
		if err != nil {
			return nil, fmt.Errorf("parsing robot: %w", err)
		}
		robots[i] = bot
	}

	return robots, nil
}

func part1(robots []robot, rs roomSize) {
	quadrants := make([]int, 4)
	for _, bot := range robots {
		quadrant := bot.quadrantAtTime(100, rs)
		if quadrant == 0 {
			continue
		}
		quadrants[quadrant-1]++
	}

	mul := 1
	for _, quad := range quadrants {
		mul *= quad
	}

	fmt.Println("Part 1:", mul)
}

func part2(robots []robot, rs roomSize) {
	t := 0
	for {
		grid := shared.NewEmptyGrid[rune](rs.width, rs.height, '.')

		for _, bot := range robots {
			pos := bot.positionAtTime(t, rs)
			grid.Set(pos, '#')
		}

		for _, row := range grid.Rows() {
			consecutiveFilled := 0
			for _, char := range row {
				if char == '#' {
					consecutiveFilled++
				} else {
					if consecutiveFilled == 31 {
						fmt.Println("Part 2:", t)
						return
					}
					consecutiveFilled = 0
				}
			}
		}
		t++
	}
}

func Run() {
	rawInput, err := shared.ReadFileByLine("days/day14/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	robots, err := parseInput(rawInput)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	rs := newRoomSize(roomWidth, roomHeight)

	part1(robots, rs)
	part2(robots, rs)
}
