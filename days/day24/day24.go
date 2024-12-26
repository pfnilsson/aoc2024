package day24

import (
	"aoc2024/shared"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type operation string

const (
	and = "AND"
	or  = "OR"
	xor = "XOR"
)

type instruction struct {
	firstKey  string
	secondKey string
	resultKey string
	operation operation
}

func newInstruction(firstKey string, secondKey string, resultKey string, operation operation) instruction {
	return instruction{firstKey, secondKey, resultKey, operation}
}

func parseState(stateRaw []string) (map[string]int, error) {
	state := make(map[string]int)
	for _, line := range stateRaw {
		lineSplit := strings.Split(line, ": ")
		key, valStr := lineSplit[0], lineSplit[1]
		val, err := strconv.Atoi(valStr)

		if err != nil {
			return nil, fmt.Errorf("error parsing state: %w", err)
		}

		state[key] = val
	}
	return state, nil
}

func parseInstructions(instructionsRaw []string) map[string]instruction {
	instructions := make(map[string]instruction)
	for _, line := range instructionsRaw {
		lineSplit := strings.Split(line, " ")
		firstKey, op, secondKey, resultKey := lineSplit[0], lineSplit[1], lineSplit[2], lineSplit[4]
		instructions[resultKey] = newInstruction(firstKey, secondKey, resultKey, operation(op))
	}
	return instructions
}

func parseInput(inputRaw [][]string) (map[string]int, map[string]instruction, error) {
	state, err := parseState(inputRaw[0])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing input: %w", err)
	}
	instructions := parseInstructions(inputRaw[1])
	return state, instructions, nil
}

func runInstruction(state map[string]int, instr instruction) error {
	firstVal, firstOk := state[instr.firstKey]
	secondVal, secondOk := state[instr.secondKey]

	if !firstOk || !secondOk {
		return errors.New("error running instruction: key not found")
	}

	switch instr.operation {
	case and:
		state[instr.resultKey] = firstVal & secondVal
	case or:
		state[instr.resultKey] = firstVal | secondVal
	case xor:
		state[instr.resultKey] = firstVal ^ secondVal
	}
	return nil
}

func runInstructions(state map[string]int, instructions map[string]instruction) {
	var instr instruction

	queue := make([]instruction, 0, len(instructions))
	for _, value := range instructions {
		queue = append(queue, value)
	}

	for len(queue) > 0 {
		instr, queue = queue[0], queue[1:]
		err := runInstruction(state, instr)
		if err != nil {
			queue = append(queue, instr)
		}
	}
}

func getIntegerFromBinary(state map[string]int, letter string) (int, error) {
	var bits []int

	for i := 0; ; i++ {
		key := fmt.Sprintf("%v%02d", letter, i)
		if val, ok := state[key]; ok {
			bits = append(bits, val)
		} else {
			break
		}
	}

	var binaryString strings.Builder
	for i := len(bits) - 1; i >= 0; i-- {
		binaryString.WriteString(strconv.Itoa(bits[i]))
	}

	decimalValue, err := strconv.ParseInt(binaryString.String(), 2, 64)
	if err != nil {
		return 0, fmt.Errorf("error calculating result: %w", err)
	}

	return int(decimalValue), nil
}

func calculateResult(state map[string]int) (int, error) {
	return getIntegerFromBinary(state, "z")
}

func findZSwaps(instructions map[string]instruction) []instruction {
	// all z are the result of XOR operations, except for these 4.
	// z45 is the most significant bit and shouldn't be swapped
	var zSwaps []instruction
	for _, instr := range instructions {
		if instr.operation != xor && instr.resultKey[0] == 'z' && instr.resultKey != "z45" {
			zSwaps = append(zSwaps, instr)
		}
	}
	return zSwaps
}

func findSwapCandidates(instructions map[string]instruction) []instruction {
	// find XOR operations that don't lead to z and don't contain x or y
	var swapCandidates []instruction
	for _, instr := range instructions {
		if instr.operation == xor && instr.resultKey[0] != 'z' && instr.firstKey[0] != 'x' && instr.firstKey[0] != 'y' && instr.secondKey[0] != 'x' && instr.secondKey[0] != 'y' {
			swapCandidates = append(swapCandidates, instr)
		}
	}
	return swapCandidates
}

func matchSwaps(zSwaps []instruction, swapCandidates []instruction, instructions map[string]instruction) map[instruction]instruction {
	matches := make(map[instruction]instruction)
	for _, zInstr := range zSwaps {
		num := zInstr.resultKey[1:]
		for _, instr := range instructions {
			if !(instr.firstKey[1:] == num && instr.secondKey[1:] == num && instr.operation == xor) {
				continue
			}

			for _, swapCandidate := range swapCandidates {
				if swapCandidate.firstKey == instr.resultKey || swapCandidate.secondKey == instr.resultKey {
					matches[zInstr] = swapCandidate
					break
				}
			}
		}
	}
	return matches
}

func resetState(state map[string]int) {
	for key := range state {
		if key[0] != 'x' && key[0] != 'y' {
			delete(state, key)
		}
	}
}

func swapInstructions(instructions map[string]instruction, swaps map[instruction]instruction) map[string]instruction {
	for swap1, swap2 := range swaps {
		key1, key2 := swap1.resultKey, swap2.resultKey

		instr1, _ := instructions[key1]
		instr2, _ := instructions[key2]

		instr1.resultKey, instr2.resultKey = instr2.resultKey, instr1.resultKey

		instructions[key1] = instr1
		instructions[key2] = instr2
	}
	return instructions
}

func toBinary(num int) string {
	if num == 0 {
		return "0"
	}

	binary := ""
	for num > 0 {
		binary = fmt.Sprintf("%d", num%2) + binary
		num = num / 2
	}
	return binary
}

func getXYZ(state map[string]int) (int, int, int, error) {
	x, err := getIntegerFromBinary(state, "x")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error getting x: %w", err)
	}
	y, err := getIntegerFromBinary(state, "y")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error getting y: %w", err)
	}
	z, err := getIntegerFromBinary(state, "z")
	if err != nil {
		return 0, 0, 0, fmt.Errorf("error getting z: %w", err)
	}
	return x, y, z, nil
}

func findWrongBitIndex(x int, y int, z int) int {
	// the last swap is the x & y bits that produced the wrong z bit
	// this function finds the index of that bit

	wrongBin := toBinary(x + y)
	actualBin := toBinary(z)

	for i := len(wrongBin) - 1; i > 0; i-- {
		if wrongBin[i] != actualBin[i] {
			return len(wrongBin) - i - 1
		}
	}
	return -1
}

func part1(state map[string]int, instructions map[string]instruction) error {
	runInstructions(state, instructions)
	res, err := calculateResult(state)
	if err != nil {
		return fmt.Errorf("error running part 1: %w", err)
	}
	fmt.Println("Part 1:", res)
	return nil
}

func part2(state map[string]int, instructions map[string]instruction) error {
	zSwaps := findZSwaps(instructions)
	swapCandidates := findSwapCandidates(instructions)
	matches := matchSwaps(zSwaps, swapCandidates, instructions)

	resetState(state)
	newInstructions := swapInstructions(instructions, matches)
	runInstructions(state, newInstructions)

	x, y, z, err := getXYZ(state)
	if err != nil {
		return fmt.Errorf("error running part 2: %w", err)
	}

	wrongBit := strconv.Itoa(findWrongBitIndex(x, y, z))

	var swaps []string
	for _, instr := range newInstructions {
		if instr.firstKey[1:] == wrongBit && instr.secondKey[1:] == wrongBit {
			swaps = append(swaps, instr.resultKey)
		}
	}

	for k, v := range matches {
		swaps = append(swaps, k.resultKey)
		swaps = append(swaps, v.resultKey)
	}

	sort.Strings(swaps)
	joinedSwaps := strings.Join(swaps, ",")

	fmt.Println("Part 2:", joinedSwaps)

	return nil
}

func Run() {
	inputRaw, err := shared.ReadFileByBlankLine("days/day24/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	state, instructions, err := parseInput(inputRaw)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	err = part1(state, instructions)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = part2(state, instructions)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

}
