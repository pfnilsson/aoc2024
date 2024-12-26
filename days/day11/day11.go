package day11

import (
	"aoc2024/shared"
	"fmt"
	"log"
)

func countDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func splitNumber(n int, numDigits int) (int, int) {
	halfLength := numDigits / 2
	divisor := 1
	for i := 0; i < halfLength; i++ {
		divisor *= 10
	}
	firstHalf := n / divisor
	secondHalf := n % divisor
	return firstHalf, secondHalf
}

func blink(num int, cache map[int][]int) []int {
	if cachedReturnValue, ok := cache[num]; ok {
		return cachedReturnValue
	}

	var returnValue []int

	if num == 0 {
		returnValue = []int{1}
		cache[num] = returnValue
		return returnValue
	}

	digitCount := countDigits(num)
	if digitCount%2 == 0 {
		first, second := splitNumber(num, digitCount)
		returnValue = []int{first, second}
		cache[num] = returnValue
		return returnValue
	}

	returnValue = []int{num * 2024}
	cache[num] = returnValue
	return returnValue
}

func parseCounts(stones []int) map[int]int {
	counts := make(map[int]int)
	for _, stone := range stones {
		counts[stone]++
	}
	return counts
}

func nextCounts(currCounts map[int]int, cache map[int][]int) map[int]int {
	newCounts := make(map[int]int)

	for num := range currCounts {
		newStones := blink(num, cache)
		for _, newStone := range newStones {
			newCounts[newStone] += currCounts[num]
		}
	}
	return newCounts
}

func part1(stoneCounts map[int]int, cache map[int][]int) map[int]int {
	for i := 0; i < 25; i++ {
		stoneCounts = nextCounts(stoneCounts, cache)
	}

	tot := 0
	for _, count := range stoneCounts {
		tot += count
	}
	fmt.Println("Part 1:", tot)

	return stoneCounts
}

func part2(stoneCounts map[int]int, cache map[int][]int) {
	for i := 0; i < 50; i++ {
		stoneCounts = nextCounts(stoneCounts, cache)
	}

	tot := 0
	for _, count := range stoneCounts {
		tot += count
	}
	fmt.Println("Part 2:", tot)
}

func Run() {
	stones, err := shared.ReadFileBySingleIntLine("days/day11/input.txt", " ")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	stoneCounts := parseCounts(stones)
	cache := make(map[int][]int)

	stoneCounts = part1(stoneCounts, cache)
	part2(stoneCounts, cache)

}
