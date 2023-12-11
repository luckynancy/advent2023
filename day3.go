package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, numberInfoList := readData()
	findPartner := func(matrix [][]int, x int, y int) (int, bool) {
		for _, row := range matrix {
			if row[1] == x && row[2] == y {
				return row[0], true
			}
		}
		return 0, false
	}

	var partI, partII int
	var starInfo [][]int

	for row, numbers := range numberInfoList {
		for _, number := range numbers {

			isSymbol, isStar, starPositionX, starPositionY := isPartNumber(row, number[1], number[2], data)
			if isSymbol || isStar {
				partI += number[0]
			}

			if isStar {
				gearPartner, isGear := findPartner(starInfo, starPositionX, starPositionY)
				if isGear {
					partII += number[0] * gearPartner
				}
				starInfo = append(starInfo, []int{number[0], starPositionX, starPositionY})
			}
		}
	}
	fmt.Println("Part I and part II are ", partI, partII)
}

func readData() ([][]string, [][][]int) {
	// open file
	file, err := os.Open("day3.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	var data [][]string
	var nrIdx [][][]int

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// find indices and number per line
		nrIdx = append(nrIdx, findNumbersAndIndices(line[0]))

		// Convert the string to a slice of characters
		char := strings.Split(line[0], "")

		data = append(data, char)
	}

	return data, nrIdx
}

func findNumbersAndIndices(input string) [][]int {

	// Define a regular expression pattern for numbers
	re := regexp.MustCompile(`\d+`)

	// Find all matches and their indices in the input string
	matchesIndex := re.FindAllStringIndex(input, -1)

	var numberInfoList [][]int

	// Extract numbers and their indices
	for _, matchIndex := range matchesIndex {
		startIndex, endIndex := matchIndex[0], matchIndex[1]
		number, _ := strconv.Atoi(input[startIndex:endIndex]) // Convert string to int
		numberInfoList = append(numberInfoList, []int{number, startIndex, endIndex})
	}

	return numberInfoList
}

func isPartNumber(row int, startIndex int, endIndex int, data [][]string) (bool, bool, int, int) {
	/*
		A number could be single digits, double digits or triple digits.
		This function check if a number is a partnumber by looking at the neighbours of a number
		given its row number, indices where the numbers starts and where it end in the datatype matrix.
	*/

	var left, right int
	var starPositionX, starPositionY int
	var isStar, isSymbol bool

	isStarOrSymbol := func(x int, y int) {
		value := data[x][y]
		switch {
		case value == "*":
			starPositionX = x
			starPositionY = y
			isStar = true
		case strings.Contains("0123456789.", value):
			//do nothing
		default:
			isSymbol = true
		}
	}

	switch startIndex {
	case 0:
		// if the first column, no left neighbour and 1 right neighbour
		left = 0
		right = 1
		isStarOrSymbol(row, endIndex)

	default:
		left = -1 // if not the first column, include the left neighbour
		isStarOrSymbol(row, startIndex-1)
		switch {
		case endIndex == len(data[0]):
			right = 0 //if last column, no right neighbour
		default:
			isStarOrSymbol(row, endIndex)
			right = 1 //if not first nor last column, include right neighbour
		}
	}

	for i := left; i < (endIndex-startIndex)+right; i++ {
		switch {
		case row == 0:
			isStarOrSymbol(row+1, startIndex+i)
		case row+1 == len(data):
			isStarOrSymbol(row-1, startIndex+i)
		default:
			isStarOrSymbol(row-1, startIndex+i)
			isStarOrSymbol(row+1, startIndex+i)
		}
	}
	return isSymbol, isStar, starPositionX, starPositionY
}
