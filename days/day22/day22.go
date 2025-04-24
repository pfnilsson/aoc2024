package day22

import (
	"fmt"
	"log"

	"aoc2024/shared"
)

func generateNextSecretNumber(num int) int {
	num = ((num * 64) ^ num) % 16777216
	num = ((num / 32) ^ num) % 16777216
	num = ((num * 2048) ^ num) % 16777216
	return num
}

func getPrice(num int) int {
	return num % 10
}

func updatePrices(prices map[string]int, num int) {
	seen := shared.NewSet[string]()
	price := getPrice(num)
	queue := shared.NewFIFOQueue[int](4)
	for range 2000 {
		num = generateNextSecretNumber(num)
		newPrice := getPrice(num)
		queue.Enqueue(newPrice - price)
		price = newPrice

		seqKey := queue.String()
		if !seen.Contains(seqKey) {
			seen.Add(seqKey)
			prices[seqKey] += newPrice
		}

	}
}

func findBest(prices map[string]int) int {
	best := 0
	for _, price := range prices {
		if price > best {
			best = price
		}
	}
	return best
}

func part1(numbers []int) {
	tot := 0
	for _, num := range numbers {
		for range 2000 {
			num = generateNextSecretNumber(num)
		}
		tot += num
	}
	fmt.Println("Part 1:", tot)
}

func part2(numbers []int) {
	prices := make(map[string]int)
	for _, num := range numbers {
		updatePrices(prices, num)
	}

	best := findBest(prices)
	fmt.Println("Part 2:", best)
}

func Run() {
	numbers, err := shared.ReadFileByLineToInt("days/day22/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	part1(numbers)
	part2(numbers)
}
