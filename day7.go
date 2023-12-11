package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Part I ", run(false))
	fmt.Println("Part II ", run(true))
}

func run(partII bool) int {

	makeEmptyMap := func() map[int]map[string]int {
		bucket := make(map[int]map[string]int)
		for i := 0; i < 7; i++ {
			bucket[i] = make(map[string]int)
		}
		return bucket
	}

	countRank := func(buckets map[int]map[string]int) int {
		count := 0
		rank := 1000
		for i := 0; i < len(buckets); i++ {
			sorted := sortMap(buckets[i])
			for j := 0; j < len(sorted); j++ {
				count += rank * sorted[j].Bid
				rank--
			}
		}
		return count
	}

	data := readData("day7.csv", partII)

	buckets := makeEmptyMap()
	for hand, bid := range data {
		switch partII {
		case true:
			temp := replaceJ(hand)
			buckets[CategorizeHand(temp)][hand] = bid
		default:
			buckets[CategorizeHand(hand)][hand] = bid
		}
	}
	return countRank(buckets)
}

func readData(str string, partII bool) map[string]int {
	// open file
	file, err := os.Open(str)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer file.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(file)

	replacementMap := map[string]string{
		//default A > K > Q > J > T > 9 > 8 > 7 > 6 > 5 > 4 > 3 > 2
		"2": "O",
		"3": "N",
		"4": "M",
		"5": "L",
		"6": "I",
		"7": "H",
		"8": "G",
		"9": "F",
		"T": "E",
		"J": "D",
		"Q": "C",
		"K": "B",
		//"A": "A", self-mapping
	}

	if partII {
		replacementMap["J"] = "P"
		//Part II: //A > K > Q > T > 9 > 8 > 7 > 6 > 5 > 4 > 3 > 2 > J
	}

	replaceValues := func(input string) string {
		//map all the values
		for key, value := range replacementMap {
			input = strings.ReplaceAll(input, key, value)
		}
		return input
	}

	hands := make(map[string]int)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row := strings.Split(line[0], " ")
		bid, _ := strconv.Atoi(row[1])
		hands[replaceValues(row[0])] = bid
	}
	return hands
}

func replaceJ(s string) string {
	//Replace the J's in str to get input to compute the right type
	counts := countOccurence(s)

	//first check the counts per char and target the most frequent one
	var maxChar rune
	maxCount := 0
	for char, count := range counts {
		if char != 'P' && count > maxCount {
			maxCount = count
			maxChar = char
		}
	}

	// need to target the char with highest values when three singles
	for char, count := range counts {
		if char != 'P' && maxCount == 1 && count == 1 && char > maxChar {
			maxChar = char
		}
	}
	return strings.ReplaceAll(s, "P", string(maxChar))
}

// CategorizeHand categorizes a poker hand based on the provided rules
func CategorizeHand(hand string) int {
	counts := countOccurence(hand)

	var freq []int
	for _, count := range counts {
		freq = append(freq, count)
	}

	// CategorizeHand categorizes a poker hand based on the frequence of cards
	switch len(freq) {
	case 1:
		return 0 // Five of a kind
	case 2:
		if freq[0] == 1 || freq[1] == 1 {
			return 1 // Four of a kind
		}
		return 2 // Full house
	case 3:
		if freq[0] == 3 || freq[1] == 3 || freq[2] == 3 {
			return 3 // Three of a kind
		}
		return 4 // Two pair
	case 4:
		return 5 //One pair
	default:
		return 6 //High card
	}
}

func countOccurence(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

// Define a struct named 'Pair'
type Pair struct {
	Hand string
	Bid  int
}

func sortMap(m map[string]int) []Pair {
	// sort the map based on the key values
	var sorted []Pair

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		sorted = append(sorted, Pair{Hand: k, Bid: m[k]})
	}
	return sorted
}
