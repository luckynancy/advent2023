package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	file, err := os.Open("day4.csv")
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	//Modify deliminater
	csvReader.Comma = '|'

	//Disable record length test in the CSV reader by setting FieldsPerRecord to a negative value.
	csvReader.FieldsPerRecord = -1

	partI := 0
	partII := 0
	copy := make(map[int]int)

	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		partII += 1 // add original

		index := strings.Index(line[0], ":")
		cardNumber, _ := strconv.Atoi(strings.Trim(line[0][5:index], " "))

		currentCopies, currentExist := copy[cardNumber]
		u := 1
		if currentExist {
			u += currentCopies
			partII += currentCopies
		}

		matches := findCommonElements(toList(line[0][index+1:]), toList(line[1]))
		//fmt.Println("cardnr", cardNumber, line[0][5:index], matches)

		points := 0
		if matches != 0 {
			//part I
			points = int(math.Pow(2, float64(matches-1)))
			//part II
			for k := 1; k < matches+1; k++ {
				if value, exists := copy[cardNumber+k]; exists {
					copy[cardNumber+k] = value + u
				} else {
					copy[cardNumber+k] = u
				}
			}
		}
		//fmt.Println("cardnr", cardNumber, matches, points)
		partI += points

	}
	fmt.Println("Part I and Part II are: ", partI, partII)
}

func findCommonElements(list1, list2 []int) int {
	commonElements := 0

	// Create a map to track elements in both lists
	elementsMap := make(map[int]bool)

	// Iterate through list1 and mark elements in the map
	for _, num := range list1 {
		elementsMap[num] = true
	}

	// Check for common elements in list2
	for _, num := range list2 {
		if elementsMap[num] {
			commonElements++
		}
	}

	return commonElements
}

func toList(str string) []int {
	//str = strings.Trim(str, "[]")

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
