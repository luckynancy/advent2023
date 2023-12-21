package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	input, startY, startX := readData("day10.csv")
	neighbours := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	partI := 0
	for _, item := range neighbours {
		distance := search(0, startX, startY, startX+item[0], startY+item[1], input)
		//fmt.Println(item[0], item[1], distance)
		partI = (map[bool]int{true: distance, false: partI})[distance > partI]

	}
	fmt.Println(partI)
}

func search(count int, currentX int, currentY int, nextX int, nextY int, input [][]string) int {
	isOutside := func(nx int, ny int) bool {
		//check the indices are within the boundary
		return (nx < 0) || (nx >= len(input)) || (ny < 0) || (ny >= len(input[0]))
	}
	if isOutside(nextX, nextY) {
		return count / 2
	}

	isStart, isGround, nX, nY := indexMap(currentX, currentY, input[nextX][nextY], nextX, nextY)
	count++

	if isStart || isGround {
		return count / 2
	}
	return search(count, nextX, nextY, nX, nY, input)
}

func indexMap(currentX int, currentY int, char string, nextX int, nextY int) (bool, bool, int, int) {
	isStart := false
	isGround := false
	var outX, outY int
	switch char {
	case "|": //vertical pipe
		outX = (map[bool]int{true: nextX + 1, false: nextX - 1})[nextX > currentX]
		outY = nextY
	case "-": //horizontal pipe
		outX = nextX
		outY = (map[bool]int{true: nextY + 1, false: nextY - 1})[nextY > currentY]
	case "L": //bend connecting north and east.
		outX = (map[bool]int{true: nextX, false: nextX - 1})[nextX > currentX]
		outY = (map[bool]int{true: nextY + 1, false: nextY})[nextX > currentX]
	case "J": //bend connecting north and west.
		outX = (map[bool]int{true: nextX, false: nextX - 1})[nextX > currentX]
		outY = (map[bool]int{true: nextY - 1, false: nextY})[nextX > currentX]
	case "7": //bend connecting south and west.
		outX = (map[bool]int{true: nextX, false: nextX + 1})[nextX < currentX]
		outY = (map[bool]int{true: nextY - 1, false: nextY})[nextX < currentX]
	case "F": //bend connecting// is a 90-degree bend connecting south and east.
		outX = (map[bool]int{true: nextX + 1, false: nextX})[nextY < currentY]
		outY = (map[bool]int{true: nextY, false: nextY + 1})[nextY < currentY]
	case "S":
		isStart = true
	default: //is ground; there is no pipe in this tile.
		isGround = true
	}
	return isStart, isGround, outX, outY
}

func readData(str string) ([][]string, int, int) {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	toStringList := func(str string) []string {
		chars := []rune(str)
		var out []string
		for i := 0; i < len(chars); i++ {
			out = append(out, string(chars[i]))
		}
		return out
	}

	var input [][]string
	var startX, startY int

	y := 0
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		strings.Split(line[0], " ")

		if strings.Contains(line[0], "S") {

			startX = strings.Index(line[0], "S")
			startY = y
		}
		input = append(input, toStringList(line[0]))
		y++
	}
	return input, startX, startY
}
