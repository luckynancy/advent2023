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
	instruction, network := readData("day8.csv")

	partI := foundTarget(0, "AAA", instruction, network, false)
	fmt.Println("Part I : ", partI)

	var steps []int
	for key := range network {
		if key[2] == 'A' {
			steps = append(steps, foundTarget(0, key, instruction, network, true))
		}
	}
	fmt.Println("Part II : ", LCM(steps[0], steps[1], steps[2:]...))
}

func readData(str string) (instr []string, network map[string][]string) {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)
	csvReader.Comma = '='
	csvReader.FieldsPerRecord = -1

	networkMap := make(map[string][]string)
	var instruction []string
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(line[0]) > 15 {
			instruction = strings.Split(line[0], "")

		} else if len(line) > 0 {

			strings.Split(line[0], "=")
			key := line[0][:3]

			strings.Trim(line[1], "(")
			nodes := strings.Split(line[1], ",")
			networkMap[key] = []string{nodes[0][2:], nodes[1][1:4]}
		}
	}
	return instruction, networkMap
}

func foundTarget(i int, key string, instruction []string, network map[string][]string, partII bool) int {
	target := (map[bool]string{true: "Z", false: "ZZZ"})[partII]

	for _, item := range instruction {
		key = (map[bool]string{true: network[key][0], false: network[key][1]})[item == "L"]
		i++
		if key == target && !partII {
			return i
		}
		if string(key[2]) == target && partII {
			return i
		}
	}
	return foundTarget(i, key, instruction, network, partII)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
