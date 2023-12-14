package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readData("day9.csv")
	partI, partII := 0, 0
	for _, item := range input {
		partI += interpolatesForwards(0, item)
		partII += interpolatesBackwards(0, true, item)
	}
	fmt.Println(partI, partII)
}

func interpolatesBackwards(k int, odd bool, intList []int) int {

	switch odd {
	case true:
		k += intList[0]
	default:
		k -= intList[0]
	}

	isLastRow, differences := computeDifferences(intList)
	if isLastRow {
		return k
	}
	return interpolatesBackwards(k, !odd, differences)

}

func interpolatesForwards(k int, intList []int) int {

	k += intList[len(intList)-1]

	isLastRow, differences := computeDifferences(intList)
	if isLastRow {
		return k
	}
	return interpolatesForwards(k, differences)

}

func computeDifferences(nums []int) (bool, []int) {

	diff := make([]int, len(nums)-1)

	count := 0
	for i := 1; i < len(nums); i++ {
		diff[i-1] = nums[i] - nums[i-1]
		if diff[i-1] == 0 {
			count++
		}
	}

	return count == len(diff), diff
}

func readData(str string) [][]int {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	var input [][]int
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		strings.Split(line[0], " ")
		input = append(input, toList(line[0]))

	}
	return input
}

func toList(str string) []int {
	// Split the string into a list of tokens
	tokens := strings.Fields(str)

	// Create a list of integers
	var intList []int

	// Convert each token to an integer and append to the list
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
		}
		intList = append(intList, num)
	}

	return intList
}
