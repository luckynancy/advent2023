package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	cmd := exec.Command("go", "build")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Start here
	data := readData("day5.csv")
	seeds, chain := cleanData(data)

	// Inner function to process rows
	findMatch := func(key int) int {
		for k := 0; k < len(chain); k++ {
			var found bool
			var target int
			for w := 0; w < len(chain[k]); w++ {
				if key >= chain[k][w][1] && key < chain[k][w][1]+chain[k][w][2] {
					target = chain[k][w][0] + (key - chain[k][w][1])
					found = true
					break
				}
			}
			key = (map[bool]int{true: target, false: key})[found]
		}
		return key
	}

	minLocation1 := math.MaxInt64
	minSeed := math.MaxInt64
	for i := 0; i < len(seeds); i++ {

		location := findMatch(seeds[i])
		if location < minLocation1 {
			minLocation1 = location
			minSeed = seeds[i]
		}
	}
	fmt.Println("Part I ", minLocation1)

	minLocation2 := math.MaxInt64

	for i := 0; i < len(seeds); i++ {
		if seeds[i] == minSeed {
			for x := 0; x < seeds[i+1]; x++ {
				location := findMatch(seeds[i] + x)
				if location < minLocation2 {
					minLocation2 = location
				}
			}
		}
	}

	fmt.Println("Part II ", minLocation2)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Compilation time: %s\n", elapsedTime)
}

func convert(str string) int {
	// Convert string to integer
	i, err := strconv.Atoi(str)

	// Check for errors during conversion
	if err != nil {
		fmt.Println("Error:", err)
	}

	return i
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

func readData(str string) []string {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	//Disable record length test in the CSV reader by setting FieldsPerRecord to a negative value.
	csvReader.FieldsPerRecord = -1

	var stringList []string
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		stringList = append(stringList, line[0])
	}

	return stringList
}

func cleanData(data []string) ([]int, map[int][][]int) {

	pass := func() {
		// pass (do nothing)
	}

	// Inner function to process rows
	processRows := func(startIndex int) [][]int {
		var intList [][]int
		for k := startIndex; ; k++ {
			if k >= len(data) || strings.Contains(data[k], "map") {
				return intList
			} else {
				intList = append(intList, toList(data[k]))
			}
		}
	}

	var seeds []int
	maps := make(map[int][][]int)

	for i := 0; i < len(data); i++ {
		item := data[i]

		switch {
		case strings.Contains(item, "seeds"):
			index := strings.Index(item, ":")
			seeds = toList(strings.Trim(item[index+1:], " "))

		case strings.Contains(item, "seed-to-soil"):
			maps[0] = processRows(i + 1)

		case strings.Contains(item, "soil-to-fertilizer"):
			maps[1] = processRows(i + 1)

		case strings.Contains(item, "fertilizer-to-water"):
			maps[2] = processRows(i + 1)

		case strings.Contains(item, "water-to-light"):
			maps[3] = processRows(i + 1)

		case strings.Contains(item, "light-to-temperature"):
			maps[4] = processRows(i + 1)

		case strings.Contains(item, "temperature-to-humidity"):
			maps[5] = processRows(i + 1)

		case strings.Contains(item, "humidity-to-location"):
			maps[6] = processRows(i + 1)

		default:
			pass()
		}
	}
	return seeds, maps

}
