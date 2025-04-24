package day23

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"aoc2024/shared"
)

func parseConnection(connectionRaw string) (string, string) {
	connectionSplit := strings.Split(connectionRaw, "-")
	return connectionSplit[0], connectionSplit[1]
}

func parseGraph(connectionsRaw []string) map[string]*shared.Set[string] {
	graph := make(map[string]*shared.Set[string])
	for _, connectionRaw := range connectionsRaw {
		from, to := parseConnection(connectionRaw)
		if _, ok := graph[from]; !ok {
			graph[from] = shared.NewSet[string]()
		}
		if _, ok := graph[to]; !ok {
			graph[to] = shared.NewSet[string]()
		}

		graph[from].Add(to)
		graph[to].Add(from)
	}
	return graph
}

func findChains(graph map[string]*shared.Set[string], start string, goal string, length int) [][]string {
	if length == 0 {
		if start == goal {
			return [][]string{{}}
		}
		return nil
	}

	var chains [][]string
	for _, node := range graph[start].Items() {
		subChains := findChains(graph, node, goal, length-1)
		for _, subChain := range subChains {
			chains = append(chains, append([]string{node}, subChain...))
		}
	}
	return chains
}

func allConnected(graph map[string]*shared.Set[string], nodes []string) bool {
	for _, node1 := range nodes {
		for _, node2 := range nodes {
			if node1 == node2 {
				continue
			}
			if !graph[node1].Contains(node2) {
				return false
			}
		}
	}
	return true
}

func bestForNode(graph map[string]*shared.Set[string], node string, currentBest int) []string {
	nodes := graph[node].Items()
	length := len(nodes)

	if length < currentBest {
		return nil
	}

	for length > currentBest {
		for _, combo := range shared.Combinations(nodes, length) {
			if allConnected(graph, combo) {
				return append(combo, node)
			}
		}
		length--
	}
	return nil
}

func part1(graph map[string]*shared.Set[string]) {
	uniqueChains := shared.NewSet[string]()
	for node := range graph {
		if node[0] != 't' {
			continue
		}

		chains := findChains(graph, node, node, 3)
		for _, chain := range chains {
			sort.Strings(chain)
			uniqueChains.Add(strings.Join(chain, ""))
		}
	}
	fmt.Println("Part 1:", uniqueChains.Size())
}

func part2(graph map[string]*shared.Set[string]) {
	var best []string
	var lengthToBeat int

	for node := range graph {
		bestCombo := bestForNode(graph, node, lengthToBeat)
		if bestCombo != nil && len(bestCombo) > lengthToBeat {
			best = bestCombo
			lengthToBeat = len(best) - 1
		}
	}
	sort.Strings(best)
	fmt.Println("Part 2:", strings.Join(best, ","))
}

func Run() {
	connectionsRaw, err := shared.ReadFileByLine("days/day23/input.txt")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	graph := parseGraph(connectionsRaw)
	part1(graph)
	part2(graph)
}
