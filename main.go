package main

import (
	"flag"
	"fmt"

	"aoc2024/days/day01"
	"aoc2024/days/day02"
	"aoc2024/days/day03"
	"aoc2024/days/day04"
	"aoc2024/days/day05"
	"aoc2024/days/day06"
	"aoc2024/days/day07"
	"aoc2024/days/day08"
	"aoc2024/days/day09"
	"aoc2024/days/day10"
	"aoc2024/days/day11"
	"aoc2024/days/day12"
	"aoc2024/days/day13"
	"aoc2024/days/day14"
	"aoc2024/days/day15"
	"aoc2024/days/day16"
	"aoc2024/days/day17"
	"aoc2024/days/day18"
	"aoc2024/days/day19"
	"aoc2024/days/day20"
	"aoc2024/days/day21"
	"aoc2024/days/day22"
	"aoc2024/days/day23"
	"aoc2024/days/day24"
	"aoc2024/days/day25"
)

func main() {
	day := flag.Int("day", 0, "Specify day number to run (1-25)")
	flag.Parse()

	switch *day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	case 8:
		day08.Run()
	case 9:
		day09.Run()
	case 10:
		day10.Run()
	case 11:
		day11.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 14:
		day14.Run()
	case 15:
		day15.Run()
	case 16:
		day16.Run()
	case 17:
		day17.Run()
	case 18:
		day18.Run()
	case 19:
		day19.Run()
	case 20:
		day20.Run()
	case 21:
		day21.Run()
	case 22:
		day22.Run()
	case 23:
		day23.Run()
	case 24:
		day24.Run()
	case 25:
		day25.Run()
	default:
		fmt.Println("Please specify a valid day (1-25).")

	}
}
