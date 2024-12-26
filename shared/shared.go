package shared

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadFileByLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadFileByLineToInt(filename string) ([]int, error) {
	lines, err := ReadFileByLine(filename)
	if err != nil {
		return nil, err
	}

	var numbers []int
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func ReadFileByBlankLine(filename string) ([][]string, error) {
	var groupedLines [][]string
	var currentGroup []string

	lines, err := ReadFileByLine(filename)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		if line == "" {
			groupedLines = append(groupedLines, currentGroup)
			currentGroup = []string{}
			continue
		}

		currentGroup = append(currentGroup, line)
	}

	groupedLines = append(groupedLines, currentGroup)
	return groupedLines, nil
}

func ReadFileByLineToSplitInts(filename string, sep string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		cerr := file.Close()
		if cerr != nil && err == nil {
			err = cerr
		}
	}(file)

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, sep)

		intLine := make([]int, len(splitLine))
		for i, num := range splitLine {
			intNum, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			intLine[i] = intNum
		}
		lines = append(lines, intLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func ReadFileBySingleIntLine(filename string, sep string) ([]int, error) {
	lines, err := ReadFileByLineToSplitInts(filename, sep)
	if err != nil {
		return nil, err
	}

	if len(lines) != 1 {
		return nil, fmt.Errorf("expected 1 line, got %d", len(lines))
	}

	return lines[0], nil
}

func ReadFileToRuneGrid(filename string) (Grid[rune], error) {
	lines, err := ReadFileByLine(filename)
	if err != nil {
		return Grid[rune]{}, err
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		row := []rune(line)
		grid[i] = row
	}
	return NewGrid(grid), nil
}

func ReadFileToRuneGridWithStartingPoint(filename string, marker rune) (Grid[rune], Point, error) {
	lines, err := ReadFileByLine(filename)
	if err != nil {
		return Grid[rune]{}, Point{}, err
	}

	var startingPoint Point

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		row := []rune(line)
		grid[i] = row
		for j, char := range row {
			if char == marker {
				startingPoint = NewPoint(j, i)
			}
		}
	}
	return NewGrid(grid), startingPoint, nil
}

func ReadFileToRuneGridWithStartingPointAndGoal(
	filename string,
	startMarker rune,
	goalMarker rune,
) (Grid[rune], Point, Point, error) {
	lines, err := ReadFileByLine(filename)
	if err != nil {
		return Grid[rune]{}, Point{}, Point{}, err
	}

	var startingPoint Point
	var goal Point

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		row := []rune(line)
		grid[i] = row
		for j, char := range row {
			if char == startMarker {
				startingPoint = NewPoint(j, i)
			}
			if char == goalMarker {
				goal = NewPoint(j, i)
			}
		}
	}
	return NewGrid(grid), startingPoint, goal, nil
}

func ReadFileToIntGrid(filename string) (Grid[int], error) {
	lines, err := ReadFileByLine(filename)
	if err != nil {
		return Grid[int]{}, err
	}

	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return Grid[int]{}, err
			}

			row[j] = num
		}
		grid[i] = row
	}
	return NewGrid(grid), nil
}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	s := &Set[T]{elements: make(map[T]struct{})}
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s *Set[T]) Add(item T) {
	s.elements[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.elements, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.elements[item]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.elements)
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.elements))
	for item := range s.elements {
		items = append(items, item)
	}
	return items
}

func (s *Set[T]) Absorb(other *Set[T]) {
	for _, item := range other.Items() {
		s.Add(item)
	}
}

func (s *Set[T]) Pop() (T, bool) {
	for item := range s.elements {
		s.Remove(item)
		return item, true
	}

	var zeroValue T
	return zeroValue, false
}

func (s *Set[T]) Peek() (T, bool) {
	for item := range s.elements {
		return item, true
	}

	var zeroValue T
	return zeroValue, false
}

func (s *Set[T]) Clone() *Set[T] {
	clone := NewSet[T]()
	clone.Absorb(s)
	return clone
}

func (s *Set[T]) Difference(other Set[T]) *Set[T] {
	diff := NewSet[T]()
	for item := range s.elements {
		if !other.Contains(item) {
			diff.Add(item)
		}
	}
	return diff
}

func (s Set[T]) String() string {
	var sb strings.Builder
	sb.WriteString("{")

	for item := range s.elements {
		sb.WriteString(fmt.Sprintf("%v, ", item))
	}

	sb.WriteString("}")
	return sb.String()
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SlicesEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Clone() Point {
	return NewPoint(p.X, p.Y)
}

func (p Point) Left() Point {
	return NewPoint(p.X-1, p.Y)
}

func (p Point) Right() Point {
	return NewPoint(p.X+1, p.Y)
}

func (p Point) Up() Point {
	return NewPoint(p.X, p.Y-1)
}

func (p Point) Down() Point {
	return NewPoint(p.X, p.Y+1)
}

func (p Point) CardinalNeighbors() []Point {
	return []Point{p.Left(), p.Right(), p.Up(), p.Down()}
}

func ManhattanDistance(p1 Point, p2 Point) int {
	return AbsInt(p1.X-p2.X) + AbsInt(p1.Y-p2.Y)
}

type Grid[T any] struct {
	Points [][]T
}

func NewGrid[T any](grid [][]T) Grid[T] {
	return Grid[T]{Points: grid}
}

func NewEmptyGrid[T any](width int, height int, fillValue T) Grid[T] {
	grid := make([][]T, height)
	for i := range grid {
		grid[i] = make([]T, width)
		for j := range grid[i] {
			grid[i][j] = fillValue
		}
	}
	return NewGrid(grid)
}

func (g *Grid[T]) Get(pt Point) T {
	return g.Points[pt.Y][pt.X]
}

func (g *Grid[T]) Set(pt Point, value T) {
	g.Points[pt.Y][pt.X] = value
}

func (g *Grid[T]) Contains(pt Point) bool {
	return pt.Y >= 0 && pt.Y < len(g.Points) && pt.X >= 0 && pt.X < len(g.Points[0])
}

func (g *Grid[T]) Clone() Grid[T] {
	newPoints := make([][]T, len(g.Points))

	for i := range g.Points {
		newRow := make([]T, len(g.Points[i]))
		copy(newRow, g.Points[i])
		newPoints[i] = newRow
	}

	return Grid[T]{Points: newPoints}
}

func (g *Grid[T]) Rows() [][]T {
	return g.Points
}

func (g *Grid[T]) Print() {
	for _, row := range g.Rows() {
		for _, item := range row {
			switch v := any(item).(type) {
			case rune:
				fmt.Print(string(v))
			default:
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func (g *Grid[T]) MaxX() int {
	return len(g.Points[0]) - 1
}

func (g *Grid[T]) MaxY() int {
	return len(g.Points) - 1
}

func UniqueSlice[T comparable](slice []T) []T {
	seen := NewSet[T]()
	uniqueSlice := make([]T, 0, len(slice))

	for _, item := range slice {
		if !seen.Contains(item) {
			uniqueSlice = append(uniqueSlice, item)
			seen.Add(item)
		}
	}
	return uniqueSlice
}

type Stack[T any] struct {
	elements []T
}

func NewStack[T any](items ...T) *Stack[T] {
	s := &Stack[T]{elements: []T{}}
	for _, item := range items {
		s.Push(item)
	}
	return s
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	maxIndex := s.Size() - 1

	if maxIndex < 0 {
		var zeroValue T
		return zeroValue, false
	}

	item := s.elements[maxIndex]
	s.elements = s.elements[:maxIndex]
	return item, true
}

type PriorityQueueEntry[T any] struct {
	Value    T
	Priority int
	index    int
}

func NewPriorityQueueEntry[T any](value T, priority int) PriorityQueueEntry[T] {
	return PriorityQueueEntry[T]{Value: value, Priority: priority}
}

type PriorityQueue[T any] []*PriorityQueueEntry[T]

func NewPriorityQueue[T any]() PriorityQueue[T] {
	return make(PriorityQueue[T], 0)
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item, ok := x.(*PriorityQueueEntry[T])
	if !ok {
		panic("Push: type assertion to *Item[T] failed")
	}
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Update(item *PriorityQueueEntry[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}

type FIFOQueue[T any] struct {
	Items    []T
	capacity int
	size     int
	front    int
	rear     int
}

func NewFIFOQueue[T any](capacity int) *FIFOQueue[T] {
	return &FIFOQueue[T]{
		Items:    make([]T, capacity),
		capacity: capacity,
		front:    0,
		rear:     0,
		size:     0,
	}
}

func (q *FIFOQueue[T]) Enqueue(item T) {
	if q.size == q.capacity {
		q.front = (q.front + 1) % q.capacity
		q.size--
	}

	q.Items[q.rear] = item
	q.rear = (q.rear + 1) % q.capacity
	q.size++
}

func (q *FIFOQueue[T]) Dequeue() T {
	if q.size == 0 {
		var zeroValue T
		return zeroValue
	}

	item := q.Items[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return item
}

func (q *FIFOQueue[T]) Front() T {
	if q.size == 0 {
		var zeroValue T
		return zeroValue
	}
	return q.Items[q.front]
}

func (q *FIFOQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *FIFOQueue[T]) Size() int {
	return q.size
}

func (q *FIFOQueue[T]) Capacity() int {
	return q.capacity
}

func (q *FIFOQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := 0; i < q.size; i++ {
		sb.WriteString(fmt.Sprintf("%v, ", q.Items[(q.front+i)%q.capacity]))
	}
	sb.WriteString("]")
	return sb.String()
}

func Combinations[T any](slice []T, r int) [][]T {
	n := len(slice)
	if r > n {
		return nil
	}

	var result [][]T

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	addCombination(slice, indices, &result)

	for {
		i := r - 1
		for i >= 0 && indices[i] == i+n-r {
			i--
		}

		if i < 0 {
			break
		}

		indices[i]++

		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}

		addCombination(slice, indices, &result)
	}

	return result
}

func addCombination[T any](slice []T, indices []int, result *[][]T) {
	var combination []T
	for _, index := range indices {
		combination = append(combination, slice[index])
	}
	*result = append(*result, combination)
}
