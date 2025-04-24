package day09

import (
	"fmt"
	"log"

	"aoc2024/shared"
)

const empty = -1

type program struct {
	id    int
	start int
	size  int
}

func newprogram(id int, start int, size int) program {
	return program{id: id, start: start, size: size}
}

type freeMemory struct {
	start int
	size  int
}

func newFreeMemory(start int, size int) freeMemory {
	return freeMemory{start: start, size: size}
}

func parseInput(rawInput string) []int {
	numbers := make([]int, len(rawInput))

	for i, char := range rawInput {
		numbers[i] = int(char - '0')
	}

	return numbers
}

func unwrapMemory(memory []int) []int {
	var unwrappedMemory []int
	var value int

	for i, block := range memory {
		if i%2 == 0 {
			value = i / 2
		} else {
			value = empty
		}

		for range block {
			unwrappedMemory = append(unwrappedMemory, value)
		}
	}
	return unwrappedMemory
}

func sortMemory(memory []int) []int {
	var currPointer int
	endPointer := len(memory) - 1

	for currPointer != endPointer {
		if memory[currPointer] != empty {
			currPointer++
			continue
		}

		if memory[endPointer] == empty {
			endPointer--
			continue
		}

		memory[currPointer], memory[endPointer] = memory[endPointer], memory[currPointer]
		currPointer++
		endPointer--
	}
	return memory
}

func parseFragMemory(memory []int) ([]program, []freeMemory) {
	var (
		free     []freeMemory
		fileID   int
		programs []program
		pointer  int
	)

	for i, block := range memory {
		if i%2 == 0 {
			fileID = i / 2
			programs = append(programs, newprogram(fileID, pointer, block))
		} else {
			free = append(free, newFreeMemory(pointer, block))
		}

		pointer += block
	}

	return programs, free
}

func sortFragMemory(programs []program, free []freeMemory) []program {
	lastProgramID := programs[len(programs)-1].id

	for i := lastProgramID; i >= 0; i-- {
		for j := range free {
			freeBlock := &free[j]

			if freeBlock.start > programs[i].start {
				break
			}

			if freeBlock.size < programs[i].size {
				continue
			}

			programs[i].start = freeBlock.start
			freeBlock.start += programs[i].size
			freeBlock.size -= programs[i].size

			break
		}
	}

	return programs
}

func part1(memory []int) {
	unwrappedMemory := unwrapMemory(memory)
	sortedMemory := sortMemory(unwrappedMemory)

	tot := 0
	for i, value := range sortedMemory {
		if value == empty {
			break
		}
		tot += i * value
	}
	fmt.Println("Part 1:", tot)
}

func part2(memory []int) {
	programs, free := parseFragMemory(memory)
	defraggedPrograms := sortFragMemory(programs, free)

	tot := 0
	for _, p := range defraggedPrograms {
		for i := range p.size {
			tot += (p.start + i) * p.id
		}
	}

	fmt.Println("Part 2:", tot)
}

func Run() {
	rawInput, err := shared.ReadFileToString("days/day09/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	memory := parseInput(rawInput)

	part1(memory)
	part2(memory)
}
