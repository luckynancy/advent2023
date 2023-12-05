package main

import (
    "encoding/csv"
    "io"
    "log"
    "os"
	"fmt"
	//"unicode"
	"strings"
	//"sort"
	"regexp"
	"strconv"
	//"reflect"
)

func main() {
    // open file
    file, err := os.Open("day2.csv")
    if err != nil {
        log.Fatal(err)
    }

    // close the file at the end of the program
    defer file.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(file)

	//Modify deliminater
	csvReader.Comma = ';'

	//Disable record length test in the CSV reader by setting FieldsPerRecord to a negative value.
	csvReader.FieldsPerRecord = -1 
	
	tot := 0
	powerSum := 0

    for {
        game, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

		var boolList []bool
		var redList, greenList, blueList []int
		var gameNumber int

        // iterate through the lines
		for i, set := range game {
			if i == 0{
				index :=strings.Index(set, ":") 
				gameNumber = convert(strings.TrimSpace(set[4:index]))
				set = set[index+1:]
			} 
			gameImpossible, red, green, blue := checkSet(set)

			redList = append(redList, red)
			greenList = append(greenList, green)
			blueList = append(blueList, blue)
			boolList = append(boolList, gameImpossible)
		}
		powerSum += findMax(redList)* findMax(greenList)* findMax(blueList)
		if !containsTrue(boolList) {
			tot += gameNumber 
		}
    }
	fmt.Println("The sum is:", tot)
	fmt.Println("The power sum is ", powerSum)
}

func checkSet(setString string)(bool, int, int, int) {

	// Split the input string into sets
	setList := strings.Split(strings.TrimSpace(setString), ";")

	// Define a regular expression to extract color and value
	regex := regexp.MustCompile(`(\d+)\s*([a-zA-Z]+)`)
	set := regex.FindAllStringSubmatch(setList[0], -1)

	limits := colorLimits()
	var boolList []bool
	var red, blue, green int

	// Process each set and check the cube value against limits
	for _, cube := range set {
		value := convert(cube[1])
		color := cube[2]
		limit := limits[color]

		if value > limit {
			boolList = append(boolList, true)
		} 

		if color=="red"{
			red = value
		} else if color == "green"{
			green = value
		} else {
			blue = value
		}
	} 
	return containsTrue(boolList), red, green, blue
}


func findMax(numbers []int) int {
    if len(numbers) == 0 {
        return 0
    }

    max := numbers[0]

    for _, num := range numbers {
        if num > max {
            max = num
        }
    }
    return max
}

func containsTrue(boolList []bool) bool {
    for _, value := range boolList {
        if value == true {
            return true
        }
    }
    return false
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

func colorLimits()map[string]int {
	// Define number of allowed cubes for each color
	colorLimits := make(map[string]int)

	colorLimits["red"]= 12 
	colorLimits["green"]= 13
	colorLimits["blue"]= 14

	return colorLimits

}