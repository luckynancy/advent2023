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
	dataType, numberInfoList := readData()
	numRows := len(numberInfoList)

	var total int
	var starInfo []NumberInfo
	for row := 0; row < numRows; row++ {
		for _, info := range numberInfoList[row] {
			isPartNumber, isStar := isPartNumber(row, info.StartIndex, info.EndIndex, dataType)
			if isPartNumber || isStar {
				total += info.Number
			}

			if isStar {
				starInfo = append(starInfo, info)
			}
		}
	}
	fmt.Println(len(starInfo))
	fmt.Println("the sum is ", total)
}

func readData() ([][]string, [][]NumberInfo) {
	// open file
	file, err := os.Open("day3.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	var dataType [][]string
	var nrIdx [][]NumberInfo

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
		var rowType []string
		for _, element := range char {
			valueType := valueType(element)
			rowType = append(rowType, valueType)
		}
		dataType = append(dataType, rowType)

	}

	return dataType, nrIdx
}

func valueType(str string) string {
	// Convert string to integer
	_, err := strconv.Atoi(str)

	var dataType string
	if err == nil {
		// if no error ocurred
		dataType = "int"
	}

	if err != nil {
		// if error when converting to int
		if str == "." {
			dataType = "period"
		} else if str == "*" {
			dataType = "star"
		} else {
			dataType = "symbol"
		}
	}

	return dataType
}

// NumberInfo represents information about a number in the string
type NumberInfo struct {
	Number     int
	StartIndex int
	EndIndex   int
}

func findNumbersAndIndices(input string) []NumberInfo {

	// Define a regular expression pattern for numbers
	re := regexp.MustCompile(`\d+`)

	// Find all matches and their indices in the input string
	matchesIndex := re.FindAllStringIndex(input, -1)

	var numberInfoList []NumberInfo

	// Extract numbers and their indices
	for _, matchIndex := range matchesIndex {
		startIndex, endIndex := matchIndex[0], matchIndex[1]
		numberStr := input[startIndex:endIndex]
		number, _ := strconv.Atoi(numberStr) // Convert string to int

		numberInfo := NumberInfo{
			Number:     number,
			StartIndex: startIndex,
			EndIndex:   endIndex,
		}

		numberInfoList = append(numberInfoList, numberInfo)
	}

	return numberInfoList
}

func isPartNumber(row int, startIndex int, endIndex int, datatype [][]string) (bool, bool) {
	numRows, numCols := len(datatype), len(datatype[0])

	var neighbours []string
	var left, right int

	if startIndex == 0 {
		// if the first column
		left = 0
		right = 1
		neighbours = append(neighbours, datatype[row][endIndex])
	} else {
		left = -1
		neighbours = append(neighbours, datatype[row][startIndex-1])

		if endIndex == numCols {
			//if last column
			right = 0
		} else {
			//if not first nor last column
			neighbours = append(neighbours, datatype[row][endIndex])
			right = 1

		}
	}

	idx := (endIndex - startIndex) + right

	for i := left; i < idx; i++ {
		if row == 0 {
			neighbours = append(neighbours, datatype[row+1][startIndex+i])
		} else if row+1 == numRows {
			neighbours = append(neighbours, datatype[row-1][startIndex+i])
		} else {
			neighbours = append(neighbours, datatype[row-1][startIndex+i])
			neighbours = append(neighbours, datatype[row+1][startIndex+i])
		}
	}

	return containsString(neighbours, "symbol"), containsString(neighbours, "star")
}

func containsString(stringList []string, target string) bool {
	for _, element := range stringList {
		if element == target {
			return true
		}
	}
	return false
}
