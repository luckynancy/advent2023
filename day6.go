package main

import (
	"fmt"
)

func main() {

	// Inner function to process races
	marginPerRace := func(item int, distance int) int {

		count := 0
		for k := 0; k < item; k++ {
			traveled := (item - k) * k
			if traveled > distance {
				count++
			}
		}
		return count
	}

	raceTime := []int{40, 82, 84, 92}
	distance := []int{233, 1011, 1110, 1487}

	margin1 := 1
	for i, item := range raceTime {
		margin1 *= marginPerRace(item, distance[i])
	}
	fmt.Println("Part I ", margin1)

	margin2 := marginPerRace(40828492, 233101111101487)
	fmt.Println("Part II ", margin2)
}
