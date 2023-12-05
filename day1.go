package main

import (
    "encoding/csv"
    "io"
    "log"
    "os"
	"fmt"
	"unicode"
	"strings"
	"sort"
)

func main() {
    // open file
    file, err := os.Open("day1.csv")
    if err != nil {
        log.Fatal(err)
    }

    // close the file at the end of the program
    defer file.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(file)

	tot := 0

    for {
        lines, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        // iterate through the lines
		for _, line := range lines {
			
			//For part II only, convert the text to numerical digits, if Part I, use line
			convertedLine := convertSubStringToNumerical(line)
			
			// Find the first and last digits
			firstDigit, lastDigit := findFirstAndLastDigits(convertedLine)
			
			// Convert to integers and concatenate
			result := concatenateDigits(firstDigit, lastDigit)
			//fmt.Println(line, convertedLine, result)
			tot += result
		}

    }
	fmt.Println("The calibration value is:", tot)
}


func findFirstAndLastDigits(s string) (firstDigit, lastDigit rune) {
	// Iterate through the string
	for _, char := range s {
		// Check if the character is a digit
		if unicode.IsDigit(char) {
			// If it's the first digit, set firstDigit
			if firstDigit == 0 {
				firstDigit = char
			}
			// Always update lastDigit
			lastDigit = char
		}
	}

	return firstDigit, lastDigit
}

func concatenateDigits(firstDigit, lastDigit rune) int {
	// Convert to integers and concatenate
	firstDigitInt := int(firstDigit - '0')
	lastDigitInt := int(lastDigit - '0')
	result := firstDigitInt*10 + lastDigitInt

	return result
}

func createNumberMap() map[string]string {
	// Define and initialize a map
    numberMap := make(map[string]string)

	// Add key-value pairs to the map
	numberMap["zero"] = "z0o"
	numberMap["one"] = "o1e"
	numberMap["two"] = "t2o"
	numberMap["three"] = "t3e"
	numberMap["four"] = "f4r"
	numberMap["five"] = "f5e"
	numberMap["six"] = "s6x"
	numberMap["seven"] = "s7n"
	numberMap["eight"] = "e8t"
	numberMap["nine"] = "n9e"
	return numberMap
}

func convertSubStringToNumerical(myString string) string {

	// save the ocurrence of value and the index into a location map
	locationMap := make(map[int]string)
	numberMap := createNumberMap()
	for key, _ := range numberMap {
		if strings.Contains(myString, key){
			location := strings.Index(myString, key)
			locationMap[location]= key
		}
	}

	// Extract keys into a slice
	var keys []int
	for key := range locationMap {
		keys = append(keys, key)
	}

	// Sort the slice of keys
	sort.Ints(keys)

	// Iterate over the sorted keys and replace the corresponding values to numerical
	for _, key := range keys {
		kii := locationMap[key]
		value := numberMap[kii]
		if strings.Contains(myString, kii){
			myString = strings.Replace(myString, kii, value, -1)
		}	
	}
	return myString
}