package main

import (
	"container/heap"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	positions, expandedRow, expandedCol, size := readData("day11.csv")
	//fmt.Println(expandedCol, expandedRow)

	partI, partII := 0, 0
	n := len(positions)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			partI += findNearestPath(size, positions[i], positions[j], expandedRow, expandedCol, false)
			partII += findNearestPath(size, positions[i], positions[j], expandedRow, expandedCol, true)
		}
	}
	fmt.Println(partI, partII)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Compilation time: %s\n", elapsedTime)
}

func readData(str string) ([][]int, map[int]struct{}, map[int]struct{}, []int) {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	var positions [][]int
	expandedRow, expandedCol := make(map[int]struct{}), make(map[int]struct{})
	row, col := 0, 0
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if row == 0 {
			col = len(line[0])
			for i := 0; i < col; i++ {
				expandedCol[i] = struct{}{}
			}
		}

		// Extract positions directly while iterating over the string
		found := false
		for column, char := range line[0] {
			if char == '#' {
				positions = append(positions, []int{row, column})
				found = true
				delete(expandedCol, column)
			}
		}

		if !found {
			expandedRow[row] = struct{}{}
		}
		row++
	}

	return positions, expandedRow, expandedCol, []int{row, col}
}

/* Finding the nearest path between two points in a 2D array using Dijkstra's algorithm.
Dijkstra's algorithm finds the shortest path between two points in a weighted graph.
In a 2D array, we can treat each cell as a node with a weight representing the cost or distance. */

// Cell represents a cell in the 2D array.
type Cell struct{ x, y, distance int }

// PriorityQueue implements a priority queue for Cells.
type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// findNearestPath finds the nearest path between two points in a 2D array.
func findNearestPath(size []int, start, end []int, wX, wY map[int]struct{}, partII bool) int {

	weight := func(x, y int) int {
		//Weight function for the distance
		_, isExpandedRow := wX[x]
		_, isExpandedCol := wY[y]
		if isExpandedRow || isExpandedCol {
			if partII {
				return 1000000
			} else {
				return 2
			}
		} else {
			return 1
		}
	}

	rows, cols := size[0], size[1]

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Directions for moving to adjacent cells (up, down, left, right)
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Cell{start[0], start[1], 0})

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*Cell)
		if curr.x == end[0] && curr.y == end[1] {
			return curr.distance
		}

		if visited[curr.x][curr.y] {
			continue //skip the rest of the loop if visited
		}

		visited[curr.x][curr.y] = true
		for i := 0; i < 4; i++ {
			nextX, nextY := curr.x+dx[i], curr.y+dy[i]

			if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols && !visited[nextX][nextY] {
				heap.Push(&pq, &Cell{nextX, nextY, curr.distance + weight(nextX, nextY)})
			}
		}
	}
	// If no path is found
	return -1
}
