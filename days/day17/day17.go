package day17

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc2024/shared"
)

type computer struct {
	registerA int
	registerB int
	registerC int
	pointer   int
}

func newComputer(a int, b int, c int) *computer {
	return &computer{registerA: a, registerB: b, registerC: c, pointer: 0}
}

func (cpu *computer) run(program []int) ([]int, error) {
	var (
		err         error
		instruction int
		operand     int
		output      []int
		outputVal   int
	)

	for cpu.pointer < len(program)-1 {
		instruction = program[cpu.pointer]
		operand = program[cpu.pointer+1]

		switch instruction {
		case 0:
			err = cpu.adv(operand)
		case 1:
			cpu.bxl(operand)
		case 2:
			err = cpu.bst(operand)
		case 3:
			cpu.jnz(operand)
		case 4:
			cpu.bxc()
		case 5:
			outputVal, err = cpu.out(operand)
			output = append(output, outputVal)
		case 6:
			err = cpu.bdv(operand)
		case 7:
			err = cpu.cdv(operand)
		}

		if err != nil {
			return nil, fmt.Errorf("error running program: %w", err)
		}
	}

	return output, nil
}

func (cpu *computer) comboOperand(operand int) (int, error) {
	if operand >= 0 && operand < 4 {
		return operand, nil
	}

	if operand == 4 {
		return cpu.registerA, nil
	}

	if operand == 5 {
		return cpu.registerB, nil
	}

	if operand == 6 {
		return cpu.registerC, nil
	}

	return 0, fmt.Errorf("invalid operand: %d", operand)
}

func (cpu *computer) adv(operand int) error {
	exp, err := cpu.comboOperand(operand)
	if err != nil {
		return fmt.Errorf("error running adv: %w", err)
	}

	cpu.registerA = cpu.registerA / twoToThePowerOf(exp)
	cpu.pointer += 2
	return nil
}

func (cpu *computer) bxl(operand int) {
	cpu.registerB = cpu.registerB ^ operand
	cpu.pointer += 2
}

func (cpu *computer) bst(operand int) error {
	val, err := cpu.comboOperand(operand)
	if err != nil {
		return fmt.Errorf("error running bst: %w", err)
	}

	cpu.registerB = val % 8
	cpu.pointer += 2
	return nil
}

func (cpu *computer) jnz(operand int) {
	if cpu.registerA != 0 {
		cpu.pointer = operand
	} else {
		cpu.pointer += 2
	}
}

func (cpu *computer) bxc() {
	cpu.registerB = cpu.registerB ^ cpu.registerC
	cpu.pointer += 2
}

func (cpu *computer) out(operand int) (int, error) {
	val, err := cpu.comboOperand(operand)
	if err != nil {
		return 0, fmt.Errorf("error running out: %w", err)
	}
	cpu.pointer += 2
	return val % 8, nil
}

func (cpu *computer) bdv(operand int) error {
	exp, err := cpu.comboOperand(operand)
	if err != nil {
		return fmt.Errorf("error running bdv: %w", err)
	}

	cpu.registerB = cpu.registerA / twoToThePowerOf(exp)
	cpu.pointer += 2
	return nil
}

func (cpu *computer) cdv(operand int) error {
	exp, err := cpu.comboOperand(operand)
	if err != nil {
		return fmt.Errorf("error running bdv: %w", err)
	}

	cpu.registerC = cpu.registerA / twoToThePowerOf(exp)
	cpu.pointer += 2
	return nil
}

func (cpu *computer) runDisassembledProgram() int {
	var err error
	var output int

	err = cpu.bst(4)
	if err != nil {
		panic("disassembled program should not error")
	}
	cpu.bxl(7)
	err = cpu.cdv(5)
	if err != nil {
		panic("disassembled program should not error")
	}
	cpu.bxc()
	cpu.bxl(4)
	output, err = cpu.out(5)
	if err != nil {
		panic("disassembled program should not error")
	}
	return output
}

func twoToThePowerOf(exp int) int {
	return 1 << exp
}

func parseComputerRegisters(registersRaw []string) (*computer, error) {
	var registers [3]int
	for i, registerRaw := range registersRaw {
		registerStr := strings.Split(registerRaw, " ")[2]
		register, err := strconv.Atoi(registerStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing computer: %w", err)
		}
		registers[i] = register
	}
	return newComputer(registers[0], registers[1], registers[2]), nil
}

func parseProgram(programRaw string) ([]int, error) {
	instructionsRaw := strings.Split(programRaw, " ")[1]
	instructionsSplit := strings.Split(instructionsRaw, ",")

	program := make([]int, len(instructionsSplit))
	for i, instr := range instructionsSplit {
		instrInt, err := strconv.Atoi(instr)
		if err != nil {
			return nil, fmt.Errorf("error parsing program: %w", err)
		}
		program[i] = instrInt
	}
	return program, nil
}

func parseInput(rawInput [][]string) (*computer, []int, error) {
	cpu, err := parseComputerRegisters(rawInput[0])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing input: %w", err)
	}

	program, err := parseProgram(rawInput[1][0])
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing input: %w", err)
	}

	return cpu, program, nil
}

func reverseEngineer(program []int, target int) int {
	programLen := len(program)
	lastIndex := programLen - 1

	if programLen == 0 {
		return target
	}

	for x := 0; x <= 7; x++ {
		cpu := newComputer(target*8+x, 0, 0)
		output := cpu.runDisassembledProgram()

		if output == program[lastIndex] {
			previous := reverseEngineer(program[:lastIndex], cpu.registerA)

			if previous == -1 {
				continue
			}
			return previous
		}
	}
	return -1
}

func part1(cpu *computer, program []int) error {
	output, err := cpu.run(program)
	if err != nil {
		return fmt.Errorf("error running part1: %w", err)
	}

	strSlice := make([]string, len(output))
	for i, value := range output {
		strSlice[i] = strconv.Itoa(value)
	}

	fmt.Println("Part 1:", strings.Join(strSlice, ","))
	return nil
}

func part2(program []int) {
	a := reverseEngineer(program, 0)
	fmt.Println("Part 2:", a)
}

func Run() {
	rawInput, err := shared.ReadFileByBlankLine("days/day17/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	cpu, program, err := parseInput(rawInput)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	err = part1(cpu, program)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part2(program)
}
